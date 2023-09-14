package listprs

import (
	"github.com/cgxarrie-go/prq/domain/ports"
	"github.com/cgxarrie-go/prq/utils"
)

// GetPRsRequest is the request to list PRs
type request struct {
	origins utils.Origins
	status  ports.PRStatus
}

// Origins implements ports.ListPRRequest.
func (r request) Origins() utils.Origins {
	return r.origins
}

// Status implements ports.ListPRRequest.
func (r request) Status() ports.PRStatus {
	return r.status
}

func NewRequest(origins utils.Origins, status ports.PRStatus) (
	req ports.ListPRRequest) {
	return request{
		origins: origins,
		status: status,
	}
}