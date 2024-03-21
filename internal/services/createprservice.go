package services

import (
	"fmt"
	"os"

	"github.com/cgxarrie-go/prq/internal/models"
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

	source, destination, err := svc.geBranches(req)
	if err != nil {
		return resp, errors.Wrap(err, "getting branches")

	}

	title := svc.getTitle(req.Title, source, destination)
	desc, err := svc.getDescription(req.Description, req.Template)
	if err != nil {
		return resp, errors.Wrap(err, "getting description")
	}

	clientReq := ports.RemoteClientCreateRequest{
		Source:      source,
		Destination: destination,
		Title:       title,
		Description: desc,
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

func (svc *createPRService) geBranches(req ports.CreatePRSvcRequest) (source, destination models.Branch, err error) {
	src, err := utils.GitCurrentBranchName()
	if err != nil {
		return source, destination,
			fmt.Errorf("getting current branch name: %w", err)
	}

	source = svc.client.Remote().NewBranch(src)

	destination = svc.client.Remote().DefaultTargetBranch()
	if req.Destination != "" {
		destination = svc.client.Remote().NewBranch(req.Destination)
	}

	return
}

func (svc createPRService) getTitle(title string, src, dest models.Branch) string {
	if title != "" {
		return title
	}

	return fmt.Sprintf("PR from %s to %s", src.Name(), dest.Name())
}

func (svc *createPRService) getDescription(desc, template string) (string, error) {

	if template == "" {
		return desc, nil
	}

	tmpl, err := os.ReadFile(template)
	if err != nil {
		tmpl = []byte("")
	}

	t := string(tmpl)
	if t == "" {
		return desc, nil
	}

	if desc == "" {
		return t, nil
	}
	return fmt.Sprintf("%s\n\n%s", desc, t), nil
}
