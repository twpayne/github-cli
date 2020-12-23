package cmd

import (
	"context"

	"github.com/google/go-github/v26/github"
	"github.com/spf13/cobra"
)

var repositoryGetCmd = &cobra.Command{
	Use:     "get",
	Args:    cobra.MinimumNArgs(1),
	Short:   "Get repositories",
	Aliases: []string{"g"},
	RunE:    makeRunE(runRepositoriesGet),
}

func init() {
	repositoriesCmd.AddCommand(repositoryGetCmd)
}

func runRepositoriesGet(ctx context.Context, client *github.Client, args []string) (interface{}, error) {
	ra, err := respositoryArgs(args)
	if err != nil {
		return nil, err
	}
	allRepositories := make([]*github.Repository, 0, len(ra))
	for _, r := range ra {
		repository, _, err := client.Repositories.Get(ctx, *r.Owner.Login, *r.Name)
		if err != nil {
			return nil, nil
		}
		allRepositories = append(allRepositories, repository)
	}
	return allRepositories, nil
}
