package cmd

import "github.com/spf13/cobra"

var (
	reposCmd = &cobra.Command{
		Use:     "repos",
		Short:   "Manage repos",
		Aliases: []string{"r"},
	}

	owner string
	repo  string
)

func init() {
	persistentFlags := reposCmd.PersistentFlags()
	persistentFlags.StringVarP(&owner, "owner", "o", "", "owner")
	persistentFlags.StringVarP(&repo, "repo", "r", "", "repo")
	rootCmd.AddCommand(reposCmd)
}
