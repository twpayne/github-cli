package cmd

import (
	"context"

	"github.com/google/go-github/v26/github"
	"github.com/spf13/cobra"
)

var (
	repositoryListReleaseAssetsCmd = &cobra.Command{
		Use:     "list-release-assets",
		Args:    cobra.NoArgs,
		Short:   "List release assets",
		Aliases: []string{"lra"},
		RunE:    makeRunE(runRepositoryListReleaseAssets),
	}

	repositoryListReleaseAssets struct {
		releaseID int64
	}
)

func init() {
	persistentFlags := repositoryListReleaseAssetsCmd.PersistentFlags()
	persistentFlags.Int64VarP(&repositoryListReleaseAssets.releaseID, "release-id", "i", 0, "release-id")
	must(repositoryListReleaseAssetsCmd.MarkPersistentFlagRequired("release-id"))
	repositoryCmd.AddCommand(repositoryListReleaseAssetsCmd)
}

func runRepositoryListReleaseAssets(ctx context.Context, client *github.Client, args []string) (interface{}, error) {
	listOptions := github.ListOptions{}
	var allReleaseAssets []*github.ReleaseAsset
	for {
		releaseAssets, response, err := client.Repositories.ListReleaseAssets(ctx, repository.owner, repository.repo, repositoryListReleaseAssets.releaseID, &listOptions)
		if err != nil {
			return nil, err
		}
		allReleaseAssets = append(allReleaseAssets, releaseAssets...)
		if response.NextPage == 0 {
			break
		}
		listOptions.Page = response.NextPage
	}
	return allReleaseAssets, nil
}
