package branch

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/AlecAivazis/survey/v2"
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

func ListAllBranches(repo string) []string {
	cmd := command.GitCommand(repo, "branch")
	var branches []string

	output, err := cmd.Output()
	catch.HandleError("Could not get list of branches", err)

	outputLines := strings.Split(string(output), "\n")
	for _, line := range outputLines {
		if line == "" {
			continue
		}
		branch := ExtractBranchOutOfOutputLine(line)
		branches = append(branches, branch)
	}

	return branches
}

func ExtractBranchOutOfOutputLine(line string) string {
	regexForBranchInsideOutput := regexp.MustCompile(`[ *]+(.*)`)
	matches := regexForBranchInsideOutput.FindStringSubmatch(line)
	branch := matches[1]
	return branch
}

const CreateNewBranchMessage = "+ Create New Branch"

func AskWhatBranchToCheckoutTo(listOfBranches []string) string {
	var destinationBranch string
	options := append(listOfBranches, CreateNewBranchMessage)
	err := survey.AskOne(
		&survey.Select{
			Message: "Select destination branch",
			Options: options,
		},
		&destinationBranch,
	)
	catch.HandleError("Could not ask what branch to checkout to", err)
	return destinationBranch
}

func AskNewBranchName() string {
	var newBranchName string
	err := survey.AskOne(
		&survey.Input{
			Message: "Choose new branch name:",
		},
		&newBranchName,
	)
	catch.HandleError("Could not ask new branch name", err)
	return newBranchName
}

func CheckoutToBranch(repo, branch string) {
	cmd := command.GitCommand(repo, "checkout", branch)
	_, err := cmd.Output()
	catch.HandleError(fmt.Sprintf("Could not checkout to the %v branch", branch), err)
}

func CreateNewBranch(repo, branchName string) {
	cmd := command.GitCommand(repo, "checkout", "-b", branchName)
	_, err := cmd.Output()
	catch.HandleError("Could create not new branch", err)
}
