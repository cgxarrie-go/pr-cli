package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func GitCurrentBranchName() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get current branch: %w", err)
	}
	return strings.TrimSpace(string(out)), nil
}

func GitCurrentOriginURL() (string, error) {
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get origin url: %w", err)
	}
	return strings.TrimSpace(string(out)), nil
}
