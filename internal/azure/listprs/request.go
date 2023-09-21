package listprs

import (
	"github.com/cgxarrie-go/prq/internal/ports"
	"github.com/cgxarrie-go/prq/internal/utils"
)

// GetPRsRequest is the request to list PRs
type request struct {
	origins utils.Remotes
}

// Origins implements ports.ListPRRequest.
func (r request) Origins() utils.Remotes {
	return r.origins
}

func NewRequest(origins utils.Remotes) (req ports.ListPRRequest) {
	return request{
		origins: origins,
	}
}