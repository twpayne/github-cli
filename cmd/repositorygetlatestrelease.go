package cmd

import (
	"context"

	"github.com/google/go-github/v26/github"
	"github.com/spf13/cobra"
)

var repositoryGetLatestReleaseCmd = &cobra.Command{
	Use:     "get-latest-release",
	Args:    cobra.NoArgs,
	Short:   "Get latest release",
	Aliases: []string{"glr"},
	RunE:    makeRunE(runRepositoryGetLatestRelease),
}

func init() {
	repositoryCmd.AddCommand(repositoryGetLatestReleaseCmd)
}

func runRepositoryGetLatestRelease(ctx context.Context, client *github.Client, args []string) (interface{}, error) {
	repositoryRelease, _, err := client.Repositories.GetLatestRelease(ctx, repository.owner, repository.repo)
	return repositoryRelease, err
}
