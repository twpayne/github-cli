package cmd

import (
	"context"
	"os"

	"github.com/google/go-github/v26/github"
	"github.com/spf13/cobra"
	"go.uber.org/multierr"
)

var (
	repositoryCreateReleaseCmd = &cobra.Command{
		Use:     "create-release [release-assets...]",
		Short:   "Create a release and optionally upload release assets",
		Aliases: []string{"cr"},
		RunE:    makeRunE(runRepositoryCreateRelease),
	}

	repositoryCreateRelease struct {
		tagName         string
		targetCommitish string
		name            string
		body            string
		draft           bool
		prerelease      bool
		releaseAssets   struct {
			label string
		}
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
	persistentFlags.StringVarP(&repositoryCreateRelease.releaseAssets.label, "release-assets-label", "l", "", "release assets label")
	repositoryCmd.AddCommand(repositoryCreateReleaseCmd)
}

func runRepositoryCreateRelease(ctx context.Context, client *github.Client, args []string) (interface{}, error) {
	release := github.RepositoryRelease{
		TagName:         &repositoryCreateRelease.tagName,
		TargetCommitish: &repositoryCreateRelease.targetCommitish,
		Name:            &repositoryCreateRelease.name,
		Body:            &repositoryCreateRelease.body,
		Draft:           &repositoryCreateRelease.draft,
		Prerelease:      &repositoryCreateRelease.prerelease,
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
	id := repositoryRelease.GetID()
	for _, name := range args {
		err = multierr.Combine(err, func() error {
			file, err := os.Open(name)
			if err != nil {
				return err
			}
			defer file.Close()
			uploadOptions := github.UploadOptions{
				Name:  name,
				Label: repositoryCreateRelease.releaseAssets.label,
			}
			releaseAsset, _, err := client.Repositories.UploadReleaseAsset(ctx, repository.owner, repository.repo, id, &uploadOptions, file)
			if err != nil {
				return err
			}
			result.ReleaseAssets = append(result.ReleaseAssets, releaseAsset)
			return nil
		}())
	}
	return &result, err
}
