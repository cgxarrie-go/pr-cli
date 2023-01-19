package azure

type GetPRsRequest struct {
	ProjectID    string
	RepositoryID string
	Status       int
}
