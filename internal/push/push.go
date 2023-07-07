package push

import (
	"fmt"
	"strings"

	"github.com/KaueSabinoSRV17/Flower/internal/catch"
	"github.com/KaueSabinoSRV17/Flower/internal/command"
)

func GetCurrentBranch(repo string) string {
	cmd := command.GitCommand(repo, "branch", "--show-current")
	output, err := cmd.Output()
	catch.HandleError("Could not get current branch", err)
	currentBranch := string(output)
	formatedBranchString := strings.Replace(currentBranch, "\n", "", 1)
	return formatedBranchString
}

func PushChanges(gitDirectory string) {
	currentBranch := GetCurrentBranch(gitDirectory)
	cmd := command.GitCommand(gitDirectory, "push", "origin", currentBranch)
	_, err := cmd.Output()
	catch.HandleError("Could not Push Changes", err)
	fmt.Printf("Pushed changes to the %v branch!", currentBranch)
}
