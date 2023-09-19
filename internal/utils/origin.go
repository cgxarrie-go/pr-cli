package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

type Origin string
type Origins []Origin

func (o Origin) IsGithub() bool {
	return strings.Contains(string(o), "github.com")
}

func (o Origin) IsAzure() bool {
	return strings.Contains(string(o), "dev.azure.com")
}

func (o Origin) IsEmpty() bool {
	return o == ""
}

func (o Origins) Append(origin Origin) Origins {
	exists := false
	for _, o := range o {
		if o == origin {
			exists = true
			break
		}
	}
	// append origin if it doesn't exist
	if !exists {
		return append(o, origin)
	}
	return o
}


func CurrentOrigin() (Origin, error) {
	return NewOrigin("")
}

func NewOrigin(path string) (Origin, error) {
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	if path != "" {
		cmd.Dir = path
	}
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get origin url: %w", err)
	}
	return Origin(strings.TrimSpace(string(out))), nil
}