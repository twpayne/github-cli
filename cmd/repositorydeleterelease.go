package cmd

import (
	"context"

	"github.com/google/go-github/v26/github"
	"github.com/spf13/cobra"
)

var (
	repositoryDeleteReleaseCmd = &cobra.Command{
		Use:     "delete-release",
		Args:    cobra.NoArgs,
		Short:   "Delete release",
		Aliases: []string{"dr"},
		RunE:    makeRunE(runRepositoryDeleteRelease),
	}

	repositoryDeleteRelease struct {
		releaseID int64
	}
)

func init() {
	persistentFlags := repositoryDeleteReleaseCmd.PersistentFlags()
	persistentFlags.Int64VarP(&repositoryDeleteRelease.releaseID, "release-id", "i", 0, "release id")
	must(repositoryDeleteReleaseCmd.MarkPersistentFlagRequired("release-id"))
	repositoryCmd.AddCommand(repositoryDeleteReleaseCmd)
}

func runRepositoryDeleteRelease(ctx context.Context, client *github.Client, args []string) (interface{}, error) {
	_, err := client.Repositories.DeleteRelease(ctx, repository.owner, repository.repo, repositoryDeleteRelease.releaseID)
	return noResult, err
}
