package use_cases

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/go-git/go-git/v5"
)

func GetRepository(pathToRepository string) *git.Worktree {

	repository, err := git.PlainOpen(pathToRepository)
	if err != nil {
		log.Fatal("Could not open Git Repo")
	}

	worktree, err := repository.Worktree()
	if err != nil {
		log.Fatal("Could not get Work Tree")
	}

	return worktree

}

func GetUnstaggedFiles() []string {

	worktree := GetRepository(".")
	status, err := worktree.Status()
	if err != nil {
		log.Fatal("Could not get Git Status")
	}

	var modifiedOrUntrackedFiles []string
	for file, s := range status {
		if s.Staging == git.Unmodified && s.Worktree == git.Unmodified {
			continue
		}
		modifiedOrUntrackedFiles = append(modifiedOrUntrackedFiles, file)
	}

	return modifiedOrUntrackedFiles

}

func AskWhatFilesToAddForStaging(files []string) []string {
	var filesToAdd []string
	prompt := &survey.MultiSelect{
		Message: "Chose the files to stage",
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

func AskCommitPrefix() string {

	var prefix string
	err := survey.AskOne(
		&survey.Select{
			Message: "Select a Prefix for the commit:",
			Options: []string{"chore", "feat", "fix", "refactor", "tests", "docs"},
		},
		&prefix,
	)

	if err != nil {
		log.Fatal("Could not ask Commit Prefix")
	}

	return prefix

}

func ResolveCommitMessage() string {

	var message string
	err := survey.AskOne(
		&survey.Input{
			Message: "What did you do? (Commit Message)",
		},
		&message,
	)

	if err != nil {
		log.Fatal("Could not ask Commit Message")
	}

	return message

}

func ConventionalCommit(prefix string, message string) {

	worktree := GetRepository(".")

	formatedMessage := fmt.Sprintf("%s: %s", prefix, message)

	worktree.Commit(formatedMessage, &git.CommitOptions{})

	fmt.Println("Sucessfully added a commit!")

}
