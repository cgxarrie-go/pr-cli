package ports

import "github.com/cgxarrie-go/prq/internal/models"

type RemoteClient interface {
	Create(req RemoteClientCreateRequest) (resp RemoteClientCreateResponse, err error)
}

type RemoteClientCreateRequest struct {
	Remote      Remote
	Source      models.Branch
	Destination models.Branch
	Title       string
	IsDraft     bool
	Description string
}

type RemoteClientCreateResponse struct {
	ID          string
	Title       string
	Description string
	URL         string
	IsDraft     bool
}
