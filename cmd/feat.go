package cmd

import (
	"github.com/KaueSabinoSRV17/Flower/internal/branch"
	"github.com/KaueSabinoSRV17/Flower/internal/command"
	"github.com/KaueSabinoSRV17/Flower/internal/feature"
	"github.com/KaueSabinoSRV17/Flower/internal/repo"
	"github.com/spf13/cobra"
)

var featCmd = &cobra.Command{
	Use:   "feat",
	Short: "checks out a new feature branch from dev",
	Long: `
	Handles unclean worktree, checks out a new feature branch from dev. 

	The new brach will be prefixed by "feature/", followed by the choosen name.

	If your Git Worktree is not clean, it will prompt you to stash your changes.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var featureName string
		repo := repo.GetRepository()
		if len(args) > 0 {
			featureName = args[0]
		} else {
			featureName = feature.AskNewFeatureName()
		}

		cleanWorkTree := feature.CheckIfWorktreeIsClean(repo)
		if !cleanWorkTree {
			stashChages := feature.AskToStashUncommitedChanges()
			if stashChages {
				command.GitCommand(repo, "stash")
			} else {
				return
			}
		} else {
			branch.CheckoutToBranch(repo, "dev")
			branch.CreateNewBranch(repo, featureName)
		}
	},
}

func init() {
	rootCmd.AddCommand(featCmd)
}
