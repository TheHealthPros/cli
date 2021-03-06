package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ListGitRemote lists all existing git remotes for a git repo
func ListGitRemote() []string {
	out, err := exec.Command("sh", "-c", "git remote").Output()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return strings.Split(string(out), "\n")
}

// AddGitRemote adds a git remote to an existing git repo
func AddGitRemote(remote string, gitURL string) {
	_, err := exec.Command("sh", "-c", fmt.Sprintf("git remote add %s %s", remote, gitURL)).Output()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

// RemoveGitRemote removes an existing git remote
func RemoveGitRemote(remote string) {
	_, err := exec.Command("sh", "-c", fmt.Sprintf("git remote remove %s", remote)).Output()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
