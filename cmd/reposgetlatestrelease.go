package cmd

import (
	"context"

	"github.com/google/go-github/v26/github"
	"github.com/spf13/cobra"
)

var (
	getLatestReleaseCmd = &cobra.Command{
		Use:     "get-latest-release",
		Args:    cobra.NoArgs,
		Short:   "Get latest release",
		Aliases: []string{"glr"},
		RunE:    makeRunE(runReposGetLatestRelease),
	}
)

func init() {
	reposCmd.AddCommand(getLatestReleaseCmd)
}

func runReposGetLatestRelease(ctx context.Context, client *github.Client, args []string) (interface{}, error) {
	rr, _, err := client.Repositories.GetLatestRelease(ctx, owner, repo)
	return rr, err
}
