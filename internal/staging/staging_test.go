package staging

import (
	"fmt"
	"testing"
)

var fileExamples = map[string]string{
	"untracked_file":        "??",
	"staged_file":           "A",
	"modified_file":         "M",
	"deleted_staged_file":   "D ",
	"deleted_unstaged_file": "D",
}

func TestGetAllFilesFromStatus(t *testing.T) {
	unstaggedFiles := GetUnstaggedFiles("/home/kaue/projects/personal/side_projects/backend/clis/git_test")
	if len(unstaggedFiles) != 15 {
		t.Error("Could not get unstagged files")
	}
	if len(unstaggedFiles[0]) <= 0 {
		t.Error("To litte name")
	}
}

func TestExtractstatusandfile(t *testing.T) {
	inputStatus := "??"
	inputFIleName := "README.md"
	statusEntry := fmt.Sprintf("%v %v", inputStatus, inputFIleName)
	status, fileName := ExtractStatusAndFile(statusEntry)
	if status != inputStatus || fileName != inputFIleName {
		t.Errorf("Could not Extract ")
	}
}

func TestIsUnstaged(t *testing.T) {
	statusEntry := fmt.Sprintf("D  ")
	if IsUnstaged(statusEntry) {
		t.Error("Should not label a staged file as an unstaged one")
	}
}

func TestIsModified(t *testing.T) {
	statusEntry := " M"
	if IsModified(statusEntry) {
		t.Error("Should not label an added file as a modified one")
	}
}
