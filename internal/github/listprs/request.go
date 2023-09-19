package listprs

import (
	"github.com/cgxarrie-go/prq/internal/ports"
	"github.com/cgxarrie-go/prq/internal/utils"
)

// GetPRsRequest is the request to list PRs
type request struct {
	origins utils.Origins
}

// Origins implements ports.ListPRRequest.
func (r request) Origins() utils.Origins {
	return r.origins
}

func NewRequest(origins utils.Origins) (req ports.ListPRRequest) {
	return request{
		origins: origins,
	}
}