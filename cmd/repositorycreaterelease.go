package cmd

import (
	"context"

	"github.com/google/go-github/v26/github"
	"github.com/spf13/cobra"
)

var (
	repositoryCreateReleaseCmd = &cobra.Command{
		Use:     "create-release [release-assets...]",
		Short:   "Create a release and optionally upload release assets",
		Aliases: []string{"cr"},
		RunE:    makeRunE(runRepositoryCreateRelease),
	}

	repositoryCreateRelease struct {
		tagName            string
		targetCommitish    string
		name               string
		body               string
		draft              bool
		prerelease         bool
		releaseAssetsLabel string
	}
)

func init() {
	persistentFlags := repositoryCreateReleaseCmd.PersistentFlags()
	persistentFlags.StringVarP(&repositoryCreateRelease.tagName, "tag-name", "t", "", "tag name")
	persistentFlags.StringVarP(&repositoryCreateRelease.targetCommitish, "target-commitish", "c", "", "target commitish")
	persistentFlags.StringVarP(&repositoryCreateRelease.name, "name", "n", "", "name")
	persistentFlags.StringVarP(&repositoryCreateRelease.body, "body", "b", "", "body")
	persistentFlags.BoolVarP(&repositoryCreateRelease.draft, "draft", "d", false, "draft")
	persistentFlags.BoolVarP(&repositoryCreateRelease.prerelease, "prerelease", "p", false, "prerelease")
	persistentFlags.StringVarP(&repositoryCreateRelease.releaseAssetsLabel, "release-assets-label", "l", "", "release assets label")
	repositoryCmd.AddCommand(repositoryCreateReleaseCmd)
}

func runRepositoryCreateRelease(ctx context.Context, client *github.Client, args []string) (interface{}, error) {
	release := github.RepositoryRelease{
		TagName:         github.String(repositoryCreateRelease.tagName),
		TargetCommitish: github.String(repositoryCreateRelease.targetCommitish),
		Name:            github.String(repositoryCreateRelease.name),
		Body:            github.String(repositoryCreateRelease.body),
		Draft:           github.Bool(repositoryCreateRelease.draft),
		Prerelease:      github.Bool(repositoryCreateRelease.prerelease),
	}
	repositoryRelease, _, err := client.Repositories.CreateRelease(ctx, repository.owner, repository.repo, &release)
	if len(args) == 0 || err != nil {
		return repositoryRelease, err
	}
	var result struct {
		RepositoryRelease *github.RepositoryRelease `json:"repositoryRelease"`
		ReleaseAssets     []*github.ReleaseAsset    `json:"releaseAssets"`
	}
	result.RepositoryRelease = repositoryRelease
	result.ReleaseAssets, err = uploadReleaseAssets(ctx, client, repositoryRelease.GetID(), repositoryCreateRelease.releaseAssetsLabel, args)
	return &result, err
}
