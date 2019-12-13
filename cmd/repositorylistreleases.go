package cmd

import (
	"context"

	"github.com/google/go-github/v26/github"
	"github.com/spf13/cobra"
)

var repositoryListReleasesCmd = &cobra.Command{
	Use:     "list-releases",
	Args:    cobra.NoArgs,
	Short:   "List releases",
	Aliases: []string{"lr"},
	RunE:    makeRunE(runRepositoryListReleases),
}

func init() {
	repositoryCmd.AddCommand(repositoryListReleasesCmd)
}

func runRepositoryListReleases(ctx context.Context, client *github.Client, args []string) (interface{}, error) {
	listOptions := github.ListOptions{}
	var allRepositoryReleases []*github.RepositoryRelease
	for {
		repositoryReleases, response, err := client.Repositories.ListReleases(ctx, repository.owner, repository.repo, &listOptions)
		if err != nil {
			return nil, err
		}
		allRepositoryReleases = append(allRepositoryReleases, repositoryReleases...)
		if response.NextPage == 0 {
			break
		}
		listOptions.Page = response.NextPage
	}
	return allRepositoryReleases, nil
}
