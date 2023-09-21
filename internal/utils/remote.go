package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

type Remote string
type Remotes []Remote

func (r Remote) IsGithub() bool {
	return strings.Contains(string(r), "github.com")
}

func (r Remote) IsAzure() bool {
	return strings.Contains(string(r), "dev.azure.com")
}

func (r Remote) IsEmpty() bool {
	return r == ""
}

func (r *Remotes) Append(origin Remote) {
	exists := false
	for _, remote := range *r {
		if remote == origin {
			exists = true
			break
		}
	}
	// append origin if it doesn't exist
	if !exists {
		*r = append(*r, origin)
	}
}


func CurrentFolderRemote() (Remote, error) {
	return FolderRemote("")
}

func FolderRemote(path string) (Remote, error) {
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	if path != "" {
		cmd.Dir = path
	}
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get remote url: %w", err)
	}
	return Remote(strings.TrimSpace(string(out))), nil
}