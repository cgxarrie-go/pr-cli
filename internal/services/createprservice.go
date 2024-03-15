package services

import (
	"fmt"
	"os"

	"github.com/cgxarrie-go/prq/internal/ports"
	"github.com/cgxarrie-go/prq/internal/utils"
	"github.com/pkg/errors"
)

type createPRService struct {
	client ports.RemoteClient
	remote ports.Remote
}

// NewCreatePRService return new instnce of github service
func NewCreatePRService(r ports.Remote, c ports.RemoteClient) ports.PRCreator {
	return createPRService{
		remote: r,
		client: c,
	}
}

// Create .
func (svc createPRService) Run(req ports.CreateSvcPRRequest) (
	resp ports.CreatePRSvcResponse, err error) {

	src, err := utils.GitCurrentBranchName()
	if err != nil {
		return resp, fmt.Errorf("getting current branch name: %w", err)
	}

	source := svc.remote.NewBranch(src)
	destination := svc.remote.DefaultTargetBranch()
	if req.Destination != "" {
		destination = svc.remote.NewBranch(req.Destination)
	}

	title := req.Title
	if req.Title == "" {
		title = fmt.Sprintf("PR from %s to %s",
			source.Name(), destination.Name())
	}

	desc := []byte("")
	if req.PRTemplate != "" {
		desc, err = os.ReadFile(req.PRTemplate)
		if err != nil {
			desc = []byte("")
		}
	}

	clientReq := ports.RemoteClientCreateRequest{
		Source:      source,
		Destination: destination,
		Title:       title,
		Description: string(desc),
		IsDraft:     req.IsDraft,
		Remote:      svc.remote,
	}

	clientResp, err := svc.client.Create(clientReq)
	if err != nil {
		return resp, errors.Wrap(err, "creating PR")
	}

	resp = ports.CreatePRSvcResponse{
		ID:          clientResp.ID,
		Title:       clientResp.Title,
		Description: clientResp.Description,
		URL:         clientResp.URL,
		IsDraft:     clientResp.IsDraft,
		Repository:  svc.remote.Repository(),
		Link:        svc.remote.PRLinkURL(clientResp.ID),
	}

	return
}
