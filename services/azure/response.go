package azure

import (
	"strconv"
	"time"

	"github.com/cgxarrie-go/prq/domain/models"
)

// GetPRsResponse response from GetPRS fro Azure Service
type GetPRsResponse struct {
	Value []GetPRsResponsePullRequest `json:"value"`
	Count int                         `json:"count"`
}

// GetPRsResponsePullRequest Pull request response item
type GetPRsResponsePullRequest struct {
	ID          int                      `json:"pullRequestId"`
	Title       string                   `json:"title"`
	Description string                   `json:"description"`
	Repo        GetPRsResponseRepository `json:"repository"`
	Status      string                   `json:"status"`
	MergeStatus string                   `json:"mergeStatus"`
	CreatedBy   GetPRsResponseUser       `json:"createdBy"`
	URL         string                   `json:"url"`
	Branches    GetPRsResponseBranches
	IsDraft     bool                 `json:"isDraft"`
	Reviewers   []GetPRsResponseUser `json:"reviewers"`
	Created     time.Time            `json:"creationDate"`
	Closed      time.Time            `json:"closedDate"`
}

// GetPRsResponseBranches pull request response branches
type GetPRsResponseBranches struct {
	Source string `json:"sourceRefName"`
	Target string `json:"targetRefName"`
}

// GetPRsResponseUser Pull request response user
type GetPRsResponseUser struct {
	DisplayName string `json:"displayName"`
	Email       string `json:"uniqueName"`
}

// GetPRsResponseRepository pull request response repository
type GetPRsResponseRepository struct {
	ID      string                `json:"id"`
	Name    string                `json:"name"`
	URL     string                `json:"url"`
	Project GetPRsResponseProject `json:"project"`
}

// GetPRsResponseProject pull request response project
type GetPRsResponseProject struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// ToPullRequest converts a GetPRsResponsePullRequest to a models.PullRequest
func (azPR GetPRsResponsePullRequest) ToPullRequest() models.PullRequest {
	return models.PullRequest{
		ID:          strconv.Itoa(azPR.ID),
		Title:       azPR.Title,
		Description: azPR.Description,
		Repository: models.Hierarchy{
			ID:   azPR.Repo.ID,
			Name: azPR.Repo.Name,
			URL:  azPR.Repo.URL,
		},
		Project: models.Hierarchy{
			ID:   azPR.Repo.Project.ID,
			Name: azPR.Repo.Project.Name,
			URL:  azPR.Repo.Project.URL,
		},
		Status:      azPR.Status,
		MergeStatus: azPR.MergeStatus,
		CreatedBy:   azPR.CreatedBy.DisplayName,
		URL:         azPR.URL,
		IsDraft:     azPR.IsDraft,
		Created:     azPR.Created,
	}
}

// CreatePRResponse .
type CreatePRResponse struct {
	ID          int                      `json:"pullRequestId"`
	Title       string                   `json:"title"`
	Description string                   `json:"description"`
	Repo        GetPRsResponseRepository `json:"repository"`
	URL         string                   `json:"url"`
	IsDraft     bool                     `json:"isDraft"`
}

func (azPR CreatePRResponse) ToPullRequest(organization string) models.CreatedPullRequest {
	return models.CreatedPullRequest{
		ID:          strconv.Itoa(azPR.ID),
		Title:       azPR.Title,
		Description: azPR.Description,
		Repository: models.Hierarchy{
			ID:   azPR.Repo.ID,
			Name: azPR.Repo.Name,
			URL:  azPR.Repo.URL,
		},
		Project: models.Hierarchy{
			ID:   azPR.Repo.Project.ID,
			Name: azPR.Repo.Project.Name,
			URL:  azPR.Repo.Project.URL,
		},
		URL:          azPR.URL,
		IsDraft:      azPR.IsDraft,
		Organization: organization,
	}
}
