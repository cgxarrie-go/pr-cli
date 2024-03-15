package remoteclient

import (
	"fmt"
	"strings"

	"github.com/cgxarrie-go/prq/internal/ports"
	"github.com/cgxarrie-go/prq/internal/remotetype"
)

func NewRemoteClient(remote string, pat string) (ports.RemoteClient, error) {

	switch true {
	case strings.Contains(remote, remotetype.Github.Name()):
		return newGhClient(pat), nil
	case strings.Contains(remote, remotetype.Azure.Name()):
		return newAzClient(pat), nil
	default:
		return nil, fmt.Errorf("remote type not supported")
	}
}
