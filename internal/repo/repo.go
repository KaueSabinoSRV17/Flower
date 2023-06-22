package repo

import (
	"log"

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
