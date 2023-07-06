package repo

import (
	"log"
	"os/exec"
	"strings"
)

func GetRepository() string {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal("Could not get Git Dir")
	}
	gitDir := string(output)
	formatedDirString := strings.Replace(gitDir, "\n", "", 1)
	return formatedDirString
}
