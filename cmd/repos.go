package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Get Github repositories",
	Long:  "Get repository information from Github",
	Run: func(cmd *cobra.Command, args []string) {
		repos := GetRepos()
		for _, repo := range repos {
			fmt.Print(repo["git_url"])
			fmt.Print(": ")
			fmt.Println(repo["description"])
		}
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
}
