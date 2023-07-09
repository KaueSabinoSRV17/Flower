package cmd

import (
	"github.com/KaueSabinoSRV17/Flower/internal/branch"
	"github.com/KaueSabinoSRV17/Flower/internal/repo"
	"github.com/spf13/cobra"
)

var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Choose a new branch to work with",
	Long: `
Choose a new branch to work with

	It list all avaliable branches. You can type to filter branch names, and press enter to checkout into one
  `,
	Run: func(cmd *cobra.Command, args []string) {
		repo := repo.GetRepository()
		branches := branch.ListAllBranches(repo)
		destinationBranch := branch.AskWhatBranchToCheckoutTo(branches)
		branch.CheckoutToBranch(repo, destinationBranch)
	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}
