package ports

import "github.com/cgxarrie-go/prq/internal/remote"

type RemoteClient interface {
	Create(req RemoteClientCreateRequest) error
}

type RemoteClientCreateRequest struct {
	remote      remote.Remote
	Source      string
	Destination string
	Title       string
	IsDraft     bool
	Description string
}
