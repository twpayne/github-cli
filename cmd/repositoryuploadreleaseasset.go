package cmd

import (
	"context"

	"github.com/google/go-github/v26/github"
	"github.com/spf13/cobra"
)

var (
	repositoryUploadReleaseAssetCmd = &cobra.Command{
		Use:     "upload-release-asset release-assets...",
		Args:    cobra.MinimumNArgs(1),
		Short:   "Upload release assets",
		Aliases: []string{"ura"},
		RunE:    makeRunE(runRepositoryUploadReleaseAsset),
	}

	repositoryUploadReleaseAsset struct {
		releaseID int64
		label     string
	}
)

func init() {
	persistentFlags := repositoryUploadReleaseAssetCmd.PersistentFlags()
	persistentFlags.Int64VarP(&repositoryUploadReleaseAsset.releaseID, "release-id", "i", 0, "release id")
	must(repositoryUploadReleaseAssetCmd.MarkPersistentFlagRequired("release-id"))
	persistentFlags.StringVarP(&repositoryUploadReleaseAsset.label, "label", "l", "", "label")
	repositoryCmd.AddCommand(repositoryUploadReleaseAssetCmd)
}

func runRepositoryUploadReleaseAsset(ctx context.Context, client *github.Client, args []string) (interface{}, error) {
	return uploadReleaseAssets(ctx, client, repositoryUploadReleaseAsset.releaseID, repositoryUploadReleaseAsset.label, args)
}
