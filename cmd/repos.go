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
			fmt.Print(repo["full_name"])
			fmt.Print(": ")
			fmt.Println(repo["description"])
			fmt.Print(": ")
			fmt.Print(repo["git_url"])
		}
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
}
