package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/v26/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var rootCmd = &cobra.Command{
	Use:           "github-cli",
	Short:         "Make GitHub API calls from the command line",
	SilenceErrors: true,
	SilenceUsage:  true,
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func makeRunE(runE func(context.Context, *github.Client, []string) (interface{}, error)) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		var httpClient *http.Client
		if accessToken, ok := os.LookupEnv("GITHUB_TOKEN"); ok {
			httpClient = oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{
				AccessToken: accessToken,
			}))
		}
		client := github.NewClient(httpClient)
		result, err := runE(ctx, client, args)
		if err != nil {
			return err
		}
		e := json.NewEncoder(os.Stdout)
		e.SetIndent("", "  ")
		return e.Encode(result)
	}
}