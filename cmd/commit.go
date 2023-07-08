/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/KaueSabinoSRV17/Flower/internal/commit"
	"github.com/KaueSabinoSRV17/Flower/internal/push"
	"github.com/KaueSabinoSRV17/Flower/internal/repo"
	"github.com/KaueSabinoSRV17/Flower/internal/staging"

	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commits always prefixing with a conventional commit",
	Long: `
  Commits always prefixing with a conventional commit

  The Conventional Commits prefixes includes "feat" for a new feature,
  "fix" for a fix and so on.
  `,
	Run: func(cmd *cobra.Command, args []string) {
		pushAfterCommit, err := cmd.Flags().GetBool("push")
		if err != nil {
			log.Fatalf("Could not get --push flag value: \n\t%v", err.Error())
		}
		repo := repo.GetRepository()
		var message string

		unstagedFiles := staging.GetUnstaggedFiles(repo)
		if len(unstagedFiles) > 0 {
			filesToStage := staging.AskWhatFilesToAddForStaging(unstagedFiles)
			staging.StageFiles(repo, filesToStage)
		}

		prefix := commit.AskCommitPrefix()

		if len(args) == 0 {
			message = commit.ResolveCommitMessage()
		} else {
			message = args[0]
		}

		commit.ConventionalCommit(repo, prefix, message)
		if pushAfterCommit {
			push.PushChanges(repo)
		}
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
	commitCmd.Flags().BoolP("push", "p", false, "Push changes after commit is completed")
}
