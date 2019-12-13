package cmd

import (
	"context"
	"os"

	"github.com/google/go-github/v26/github"
	"github.com/spf13/cobra"
	"go.uber.org/multierr"
)

var (
	repositoryCmd = &cobra.Command{
		Use:     "repository",
		Short:   "Manage repositories",
		Aliases: []string{"r", "repo"},
	}

	repository struct {
		owner string
		repo  string
	}
)

func init() {
	persistentFlags := repositoryCmd.PersistentFlags()
	persistentFlags.StringVarP(&repository.owner, "owner", "o", "", "owner")
	must(repositoryCmd.MarkPersistentFlagRequired("owner"))
	persistentFlags.StringVarP(&repository.repo, "repo", "r", "", "repo")
	must(repositoryCmd.MarkPersistentFlagRequired("repo"))
	rootCmd.AddCommand(repositoryCmd)
}

func uploadReleaseAsset(ctx context.Context, client *github.Client, releaseID int64, label, arg string) (*github.ReleaseAsset, error) {
	file, err := os.Open(arg)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	uploadOptions := github.UploadOptions{
		Name:  arg,
		Label: label,
	}
	releaseAsset, _, err := client.Repositories.UploadReleaseAsset(ctx, repository.owner, repository.repo, releaseID, &uploadOptions, file)
	return releaseAsset, err
}

func uploadReleaseAssets(ctx context.Context, client *github.Client, releaseID int64, label string, args []string) ([]*github.ReleaseAsset, error) {
	allReleaseAssets := make([]*github.ReleaseAsset, 0, len(args))
	var allErrors error
	for _, arg := range args {
		releaseAsset, err := uploadReleaseAsset(ctx, client, releaseID, label, arg)
		allReleaseAssets = append(allReleaseAssets, releaseAsset)
		allErrors = multierr.Combine(allErrors, err)
	}
	return allReleaseAssets, allErrors
}
