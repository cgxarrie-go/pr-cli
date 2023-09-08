package azure

import "github.com/cgxarrie-go/prq/services/azure/status"

// GetPRsRequest is the request for Azure service
type GetPRsRequest struct {
	// ProjectRepos is a list of project repositories
	// Key is ProjectID, value is slice of RepositoryID
	ProjectRepos map[string][]string
	Status       status.Status
}

type CreatePRRequest struct {
	Destination string
	Title       string
}
