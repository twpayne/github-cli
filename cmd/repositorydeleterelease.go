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
		id int64
	}
)

func init() {
	persistentFlags := repositoryDeleteReleaseCmd.PersistentFlags()
	persistentFlags.Int64VarP(&repositoryDeleteRelease.id, "id", "i", 0, "id")
	repositoryCmd.AddCommand(repositoryDeleteReleaseCmd)
}

func runRepositoryDeleteRelease(ctx context.Context, client *github.Client, args []string) (interface{}, error) {
	_, err := client.Repositories.DeleteRelease(ctx, repository.owner, repository.repo, repositoryDeleteRelease.id)
	return noResult, err
}
