package use_cases

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
)

func ConventionalCommit() {

	r, err := git.PlainOpen("/home/kaue/projects/personal/backend/clis/flower")
	if err != nil {
		log.Fatal("Could not open Git Repo")
	}

	w, err := r.Worktree()
	if err != nil {
		log.Fatal("Could not get Work Tree")
	}

	w.Commit("chore: commit with fixed conventional commit prefix", &git.CommitOptions{})

	fmt.Println("Commitado")

}
