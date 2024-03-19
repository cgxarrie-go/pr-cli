package services

import (
	"github.com/cgxarrie-go/prq/internal/ports"
	"github.com/pkg/errors"
)

type getPRsService struct {
	client ports.RemoteClient
}

func NewGetPRsService(c ports.RemoteClient) ports.PRReader {
	return getPRsService{
		client: c,
	}
}

func (svc getPRsService) Run() (resp ports.GetPRsSvcResponse) {

	resp.Remote = svc.client.Remote().Path()

	clResp, err := svc.client.Get()
	if err != nil {
		resp.Error = errors.Wrap(err, "getting prs")
	}

	resp.Count = len(clResp)
	resp.PullRequests = make([]ports.GetPRsSvcResponseItem, len(clResp))
	for i, pr := range clResp {
		resp.PullRequests[i] = ports.GetPRsSvcResponseItem{
			ID:          pr.ID,
			Title:       pr.Title,
			Description: pr.Description,
			IsDraft:     pr.IsDraft,
			Link:        svc.client.Remote().PRLinkURL(pr.ID),
			Created:     pr.Created,
			CreatedBy:   pr.CreatedBy,
		}
	}

	return resp
}
