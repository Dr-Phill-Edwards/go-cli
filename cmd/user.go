package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Get Github user",
	Long:  "Get authenticated user information from Github",
	Run: func(cmd *cobra.Command, args []string) {
		user := GetUser()
		fmt.Print(user["login"])
		fmt.Print(": ")
		fmt.Println(user["name"])
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
}
