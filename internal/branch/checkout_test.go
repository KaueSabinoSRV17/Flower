package branch

import (
	"fmt"
	"testing"

	"github.com/KaueSabinoSRV17/Flower/internal/repo"
)

func TestExtractBranchOutOfOutputLine(t *testing.T) {
	branch := "current_branch"
	branchesFromOutput := fmt.Sprintf("* %v", branch)
	onlyBranch := ExtractBranchOutOfOutputLine(branchesFromOutput)
	if onlyBranch != branch {
		t.Error("Could extract Branch out of Output Line")
	}
}

func TestListAllBranches(t *testing.T) {
	currentGitRepo := repo.GetRepository()
	currentBranch := GetCurrentBranch(currentGitRepo)
	expectedBranches := []string{"dev", currentBranch, "main"}
	Resultedbranches := ListAllBranches(currentGitRepo)
	if len(Resultedbranches) != 3 {
		t.Error("Less Branches than Expected")
	}
	for index, branch := range expectedBranches {
		if branch != Resultedbranches[index] {
			t.Errorf("Missing %v branch", branch)
		}
	}
}
