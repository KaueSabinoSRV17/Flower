/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/KaueSabinoSRV17/Flower/use_cases"

	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commits always prefixing with a conventional commit",
	Long: `
  Commits always prefixing with a conventional commit

  The Conventional Commits prefixes includes "feat" for a new feature,
  "fix" for a fix and so on.
  `,
	Run: func(cmd *cobra.Command, args []string) {
		use_cases.ConventionalCommit()
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
