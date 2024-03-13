package listprs

import (
	"github.com/cgxarrie-go/prq/internal/ports"
	"github.com/cgxarrie-go/prq/internal/remote"
)

// GetPRsRequest is the request to list PRs
type request struct {
	origins remote.Remotes
}

// Origins implements ports.ListPRRequest.
func (r request) Origins() remote.Remotes {
	return r.origins
}

func NewRequest(origins remote.Remotes) (req ports.ListPRRequest) {
	return request{
		origins: origins,
	}
}
