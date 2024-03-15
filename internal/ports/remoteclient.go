package ports

import (
	"time"

	"github.com/cgxarrie-go/prq/internal/models"
)

type RemoteClient interface {
	Create(req RemoteClientCreateRequest) (resp RemoteClientCreateResponse, err error)
	Get(req RemoteClientGetRequest) (resp []RemoteClientGetResponse, err error)
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

type RemoteClientGetRequest struct {
	Remote Remote
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
