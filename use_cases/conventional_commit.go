package use_cases

import (
	"fmt"
	"log"

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

func ConventionalCommit() {

	worktree := getRepository("/home/kaue/projects/personal/backend/clis/flower")

	worktree.Commit("refactor: separing concerns on git operations", &git.CommitOptions{})

	fmt.Println("Commitado")

}
