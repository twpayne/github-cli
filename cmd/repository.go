package cmd

import "github.com/spf13/cobra"

var (
	repositoryCmd = &cobra.Command{
		Use:     "repository",
		Short:   "Manage repositories",
		Aliases: []string{"r", "repo"},
	}

	owner string
	repo  string
)

func init() {
	persistentFlags := repositoryCmd.PersistentFlags()
	persistentFlags.StringVarP(&owner, "owner", "o", "", "owner")
	persistentFlags.StringVarP(&repo, "repo", "r", "", "repo")
	rootCmd.AddCommand(repositoryCmd)
}
