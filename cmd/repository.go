package cmd

import "github.com/spf13/cobra"

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
	rootCmd.MarkFlagRequired("owner")
	persistentFlags.StringVarP(&repository.repo, "repo", "r", "", "repo")
	rootCmd.MarkFlagRequired("repo")
	rootCmd.AddCommand(repositoryCmd)
}
