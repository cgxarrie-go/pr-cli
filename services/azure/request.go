package azure

import (
	"github.com/cgxarrie-go/prq/services/azure/status"
	"github.com/cgxarrie-go/prq/utils"
)

// GetPRsRequest is the request for Azure service
type GetPRsRequest struct {
	// ProjectRepos is a list of project repositories
	// Key is ProjectID, value is slice of RepositoryID
	Origins utils.Origins
	Status  status.Status
}

type CreatePRRequest struct {
	Destination string
	Title       string
}
