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
}

// NewCreatePRService return new instnce of github service
func NewCreatePRService(c ports.RemoteClient) ports.PRCreator {
	return createPRService{
		client: c,
	}
}

func (svc createPRService) Run(req ports.CreatePRSvcRequest) (
	resp ports.CreatePRSvcResponse, err error) {

	src, err := utils.GitCurrentBranchName()
	if err != nil {
		return resp, fmt.Errorf("getting current branch name: %w", err)
	}

	source := svc.client.Remote().NewBranch(src)
	destination := svc.client.Remote().DefaultTargetBranch()
	if req.Destination != "" {
		destination = svc.client.Remote().NewBranch(req.Destination)
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
		Repository:  svc.client.Remote().Repository(),
		Link:        svc.client.Remote().PRLinkURL(clientResp.ID),
	}

	return
}
