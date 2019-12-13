package cmd

import (
	"context"
	"os"

	"github.com/google/go-github/v26/github"
	"github.com/spf13/cobra"
)

var (
	repositoryUploadReleaseAssetCmd = &cobra.Command{
		Use:     "upload-release-asset",
		Args:    cobra.ExactArgs(1),
		Short:   "Upload a release asset",
		Aliases: []string{"ura"},
		RunE:    makeRunE(runRepositoryUploadReleaseAsset),
	}

	repositoryUploadReleaseAsset struct {
		release struct {
			id int64
		}
		name  string
		label string
	}
)

func init() {
	persistentFlags := repositoryUploadReleaseAssetCmd.PersistentFlags()
	persistentFlags.Int64VarP(&repositoryUploadReleaseAsset.release.id, "release-id", "i", 0, "release id")
	repositoryUploadReleaseAssetCmd.MarkFlagRequired("release-id")
	persistentFlags.StringVarP(&repositoryUploadReleaseAsset.name, "name", "n", "", "name")
	persistentFlags.StringVarP(&repositoryUploadReleaseAsset.label, "label", "l", "", "label")
	repositoryCmd.AddCommand(repositoryUploadReleaseAssetCmd)
}

func runRepositoryUploadReleaseAsset(ctx context.Context, client *github.Client, args []string) (interface{}, error) {
	file, err := os.Open(args[0])
	if err != nil {
		return nil, err
	}
	defer file.Close()
	uploadOptions := github.UploadOptions{
		Name:  args[0],
		Label: repositoryUploadReleaseAsset.label,
	}
	if repositoryUploadReleaseAsset.name != "" {
		uploadOptions.Name = repositoryUploadReleaseAsset.name
	}
	releaseAsset, _, err := client.Repositories.UploadReleaseAsset(ctx, repository.owner, repository.repo, repositoryUploadReleaseAsset.release.id, &uploadOptions, file)
	return releaseAsset, err
}
