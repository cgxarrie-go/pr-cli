package remoteclient

import (
	"github.com/cgxarrie-go/prq/internal/config"
	"github.com/cgxarrie-go/prq/internal/errors"
	"github.com/cgxarrie-go/prq/internal/ports"
	"github.com/cgxarrie-go/prq/internal/remotetype"
)

func NewRemoteClient(remote ports.Remote) (ports.RemoteClient, error) {

	switch remote.Type() {
	case remotetype.Github:
		return newGhClient(remote, config.GetInstance().Github.PAT), nil
	case remotetype.Azure:
		return newAzClient(remote, config.GetInstance().Azure.PAT), nil
	default:
		return nil, errors.NewUnknownRepositoryType(remote.Type().Name())
	}
}
