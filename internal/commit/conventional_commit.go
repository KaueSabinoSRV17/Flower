package commit

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/KaueSabinoSRV17/Flower/internal/command"
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

func ConventionalCommit(gitDirectory, prefix, message string) {
	formatedMessage := fmt.Sprintf(`%s: %s`, prefix, message)
	cmd := command.GitCommand(gitDirectory, "commit", "-m", formatedMessage)
	_, err := cmd.Output()
	if err != nil {
		log.Fatalf("Could not get Commit\n\t%v", err.Error())
	}
	fmt.Println("Sucessfully added a commit!")
}
