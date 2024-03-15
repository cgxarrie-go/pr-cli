package services

import (
	"github.com/cgxarrie-go/prq/internal/ports"
	"github.com/cgxarrie-go/prq/internal/remote"
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

func (svc getPRsService) Run() (resp ports.GetPRsSvcResponse, err error) {

	resp.Remote = r

	rem, err := remote.NewRemote(r)
	if err != nil {
		resp.Error = errors.Wrap(err, "getting remote")
	}

	clReq := ports.RemoteClientGetRequest{
		Remote: rem,
	}
	clResp, err := svc.client.Get(clReq)
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
			Link:        rem.PRLinkURL(pr.ID),
		}
	}

	return resp, nil
}
