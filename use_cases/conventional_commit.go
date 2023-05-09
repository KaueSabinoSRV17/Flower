package use_cases

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/go-git/go-git/v5"
)

func getRepository(pathToRepository string) *git.Worktree {

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

func AskCommitPrefix() string {

	var prefix string
	err := survey.AskOne(
		&survey.Select{
			Message: "Select a Prefix for the commit:",
			Options: []string{"chore", "feat", "fix", "refactor"},
		},
		&prefix,
	)

	if err != nil {
		log.Fatal("Could not as Commit Prefix")
	}

	return prefix

}

func ResolveCommitMessage(possibleCommitMessage string) string {
	if possibleCommitMessage == "" {
		err := survey.AskOne(
			&survey.Input{
				Message: "What did you do? (Will be the Commit message)",
			},
			&possibleCommitMessage)
		if err != nil {
			log.Fatal("Could not ask Commit Message")
		}
	}
	return possibleCommitMessage
}

func ConventionalCommit(prefix string, message string) {

	worktree := getRepository(".")

	formatedMessage := fmt.Sprintf("%s: %s", prefix, message)

	worktree.Commit(formatedMessage, &git.CommitOptions{})

	fmt.Println("Sucessfully added a commit!")

}
