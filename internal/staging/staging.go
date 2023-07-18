package staging

import (
	"log"
	"regexp"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/KaueSabinoSRV17/Flower/internal/command"
)

func GetUnstaggedFiles(dir string) []string {
	files := GetAllFilesFromStatus(dir)
	var unstagedFiles []string
	for _, file := range files {
		if file == "" {
			continue // There is always an empty string in the list. If we reached and dont continue the loop, we'll have some problems
		}
		status, fileName := ExtractStatusAndFile(file)
		if IsUnstaged(status) {
			unstagedFiles = append(unstagedFiles, fileName)
		}
	}
	return unstagedFiles
}

func GetAllFilesFromStatus(dir string) []string {
	cmd := command.GitCommand(dir, "status", "-s")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Could not get Status %v\n\t", err.Error())
	}
	files := string(output)
	arrayOfFiles := strings.Split(files, "\n")
	return arrayOfFiles
}

func ExtractStatusAndFile(stageEntry string) (status, fileName string) {
	regex := regexp.MustCompile("([A-Z ?]+) (.*)")
	matches := regex.FindStringSubmatch(stageEntry)
	if len(matches) >= 3 {
		status = matches[1]
		fileName = matches[2]
	} else {
		log.Fatal("Could not Extract Status and Files")
	}
	return
}

func IsUnstaged(status string) bool {
	alreadyStagedStatus := []string{"A", "D  "}
	for _, statusName := range alreadyStagedStatus {
		if status == statusName {
			return false
		}
	}
	return true
}

func IsModified(status string) bool {
	modifiedStatus := " M"
	return status == modifiedStatus
}

func StageFiles(dir string, files []string) {
	args := append([]string{"add"}, files...)
	cmd := command.GitCommand(dir, args...)
	_, err := cmd.Output()
	if err != nil {
		log.Fatal("Could not add Files")
	}
}

func StashChanges(dir string) {
	cmd := command.GitCommand(dir, "stash")
	_, err := cmd.Output()
	if err != nil {
		log.Fatal("Could not stash changes")
	}
}

func AskWhatFilesToAddForStaging(files []string) []string {
	var filesToAdd []string
	prompt := &survey.MultiSelect{
		Message: "Chose the files to stage:",
		Options: files,
	}
	err := survey.AskOne(prompt, &filesToAdd)
	if err != nil {
		log.Fatal("Could not Ask Files to Stage")
	}
	return filesToAdd
}
