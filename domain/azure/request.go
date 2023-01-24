package azure

// GetPRsRequest is the request for Azure service
type GetPRsRequest struct {
	// ProjectRepos is a list of project repositories
	// Key is ProjectID, value is slice of RepositoryID
	ProjectRepos map[string][]string
	Status       int
}
