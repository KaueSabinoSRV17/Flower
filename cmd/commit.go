/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/KaueSabinoSRV17/Flower/internal/commit"
	"github.com/KaueSabinoSRV17/Flower/internal/repo"
	"github.com/KaueSabinoSRV17/Flower/internal/staging"

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
		repo := repo.GetRepository(".")
		var message string

		unstagedFiles := staging.GetUnstaggedFiles(repo)
		if len(unstagedFiles) > 0 {
			filesToStage := staging.AskWhatFilesToAddForStaging(unstagedFiles)
			go staging.StageFiles(filesToStage, repo)
		}

		prefix := commit.AskCommitPrefix()

		if len(args) == 0 {
			message = commit.ResolveCommitMessage()
		} else {
			message = args[0]
		}

		commit.ConventionalCommit(prefix, message, repo)
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
