package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/cgxarrie-go/prq/internal/ports"
	"github.com/cgxarrie-go/prq/internal/remote"
	"github.com/pkg/errors"
)

func GitCurrentBranchName() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get current branch: %w", err)
	}
	return strings.TrimSpace(string(out)), nil
}

func IsGitRepo(path string) bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	cmd.Dir = path
	out, err := cmd.Output()
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(out)) == "true"
}

func CurrentFolderRemote() (ports.Remote, error) {
	return folderRemote("")
}

func CurrentFolderTreeRemotes() (remotes remote.Remotes, err error) {
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

func listRemotes(root string) (remotes remote.Remotes, err error) {

	// Define a function to be called for each directory and subdirectory
	visit := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("\r\033[K%s/%s", path, info.Name())
		if info.IsDir() {
			remote, err := folderRemote(path)
			if err != nil {
				return nil
			}
			remotes[remote] = struct{}{}
			fmt.Printf("\r\033[K%s\n", remote)
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

func folderRemote(path string) (ports.Remote, error) {
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	if path != "" {
		cmd.Dir = path
	}
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get remote url: %w", err)
	}
	return remote.NewRemote(string(out))
}
