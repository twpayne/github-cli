package cmd

import (
	"fmt"
	"strings"

	"github.com/google/go-github/v26/github"
	"github.com/spf13/cobra"
)

var repositoriesCmd = &cobra.Command{
	Use:     "repositories",
	Short:   "Manage repositories",
	Aliases: []string{"rs", "repos"},
}

func init() {
	rootCmd.AddCommand(repositoriesCmd)
}

func respositoryArgs(args []string) ([]*github.Repository, error) {
	result := make([]*github.Repository, 0, len(args))
	for _, arg := range args {
		elems := strings.SplitN(arg, "/", 2)
		if len(elems) != 2 || elems[0] == "" || elems[1] == "" {
			return nil, fmt.Errorf("%s: cannot determine owner and repository", arg)
		}
		result = append(result, &github.Repository{
			Owner: &github.User{
				Login: &elems[0],
			},
			Name: &elems[1],
		})
	}
	return result, nil
}
