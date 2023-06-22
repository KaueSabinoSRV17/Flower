package commit

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/go-git/go-git/v5"
)

func AskCommitPrefix() string {
	var prefix string
	err := survey.AskOne(
		&survey.Select{
			Message: "Select a Prefix for the commit:",
			Options: []string{"chore", "feat", "fix", "refactor", "tests", "docs", "build"},
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
			Message: "What did you do? (Commit Message):",
		},
		&message,
	)

	if err != nil {
		log.Fatal("Could not ask Commit Message")
	}

	return message
}

func ConventionalCommit(prefix string, message string, worktree *git.Worktree) {
	formatedMessage := fmt.Sprintf("%s: %s", prefix, message)
	worktree.Commit(formatedMessage, &git.CommitOptions{})
	fmt.Println("Sucessfully added a commit!")
}
