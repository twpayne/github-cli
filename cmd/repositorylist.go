package cmd

import (
	"context"

	"github.com/google/go-github/v26/github"
	"github.com/spf13/cobra"
)

var (
	repositoryListCmd = &cobra.Command{
		Use:     "list",
		Args:    cobra.NoArgs,
		Short:   "List repositories",
		Aliases: []string{"l"},
		RunE:    makeRunE(runRepositoryList),
	}

	repositoryList struct {
		user        string
		visibility  string
		affiliation string
		_type       string
		sort        string
		direction   string
	}
)

func init() {
	persistentFlags := repositoryListCmd.PersistentFlags()
	persistentFlags.StringVarP(&repositoryList.user, "user", "u", "", "user")
	persistentFlags.StringVarP(&repositoryList.visibility, "visibility", "v", "", "visibility")
	persistentFlags.StringVarP(&repositoryList.affiliation, "affiliation", "a", "", "affiliation")
	persistentFlags.StringVarP(&repositoryList._type, "type", "t", "", "type")
	persistentFlags.StringVarP(&repositoryList.sort, "sort", "s", "", "sort")
	persistentFlags.StringVarP(&repositoryList.direction, "direction", "d", "", "direction")
	repositoryCmd.AddCommand(repositoryListCmd)
}

func runRepositoryList(ctx context.Context, client *github.Client, args []string) (interface{}, error) {
	repositoryListOptions := github.RepositoryListOptions{
		Visibility:  repositoryList.visibility,
		Affiliation: repositoryList.affiliation,
		Type:        repositoryList._type,
		Sort:        repositoryList.sort,
		Direction:   repositoryList.direction,
	}
	var allRepositories []*github.Repository
	for {
		repositories, response, err := client.Repositories.List(ctx, repositoryList.user, &repositoryListOptions)
		if err != nil {
			return nil, err
		}
		allRepositories = append(allRepositories, repositories...)
		if response.NextPage == 0 {
			break
		}
		repositoryListOptions.Page = response.NextPage
	}
	return allRepositories, nil
}
