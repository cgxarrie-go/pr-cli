package ports

import (
	"time"

	"github.com/cgxarrie-go/prq/internal/models"
)

type RemoteClient interface {
	Create(req RemoteClientCreateRequest) (resp RemoteClientCreateResponse, err error)
	Get() (resp []RemoteClientGetResponse, err error)
	Remote() Remote
	Open(id string) error
	OpenCode() error
}

type RemoteClientCreateRequest struct {
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

type RemoteClientGetResponse struct {
	ID          string
	Title       string
	Description string
	Status      string
	CreatedBy   string
	IsDraft     bool
	Created     time.Time
}
