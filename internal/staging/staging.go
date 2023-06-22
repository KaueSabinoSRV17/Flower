package staging

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/go-git/go-git/v5"
)

func GetUnstaggedFiles(worktree *git.Worktree) []string {
	status, err := worktree.Status()
	if err != nil {
		log.Fatal("Could not get Git Status")
	}

	var modifiedOrUntrackedFiles []string
	for file, s := range status {
		if s.Worktree == git.Modified || s.Worktree == git.Untracked {
			modifiedOrUntrackedFiles = append(modifiedOrUntrackedFiles, file)
		}
	}

	return modifiedOrUntrackedFiles
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

func StageFiles(files []string, worktree *git.Worktree) {
	for _, file := range files {
		_, err := worktree.Add(file)
		if err != nil {
			log.Fatal("Could not add " + file + " file")
		}
	}
}
