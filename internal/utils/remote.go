package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
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

func CurrentFolderTreeRemotes() (remotes Remotes, err error) {
    currentDir, err := os.Getwd()
    if err != nil {
        return remotes,
		 errors.Wrapf(err, "getting current directory")
    }

    remotes, err = listRemotes(currentDir)
    if err != nil {
        return remotes, 
		errors.Wrapf(err, "walking directories to find origins")
    }

	return
}

func listRemotes(root string) (remotes Remotes, err error) {

    // Define a function to be called for each directory and subdirectory
    visit := func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() {
			remote, err := FolderRemote(path)
            if err != nil {
				return nil
			}
			remotes = append(remotes, remote)
			return filepath.SkipDir
        }
        return nil
    }

    // Start walking the directory tree
    err = filepath.Walk(root, visit)
    if err != nil {
        return nil, errors.Wrapf(err, "getting remotes from directory tree")
    }

    return 
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