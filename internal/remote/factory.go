package remote

import (
	"fmt"
	"strings"

	"github.com/cgxarrie-go/prq/internal/ports"
	"github.com/cgxarrie-go/prq/internal/remotetype"
)

func NewRemote(r string) (ports.Remote, error) {

	switch true {
	case strings.Contains(r, remotetype.Github.Name()):
		return newGithubRemote(r), nil
	case strings.Contains(r, remotetype.Azure.Name()):
		return NewAzureRemote(r), nil
	default:
		return nil, fmt.Errorf("remote type not supported")
	}
}
