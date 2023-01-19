package azure

import (
	"time"

	"github.com/cgxarrie/pr-go/domain/models"
)

type GetPRsResponse struct {
	Value []GetPRsResponsePullRequest `json:"value"`
	Count int                         `json:"count"`
}

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
	IsDraft     bool `json:"isDraft"`
	Dates       GetPRsResponseDates
	Reviewers   []GetPRsResponseUser `json:"reviewers"`
}

type GetPRsResponseDates struct {
	Created time.Time `json:"creationDate"`
	Closed  time.Time `json:"closedDate"`
}
type GetPRsResponseBranches struct {
	Source string `json:"sourceRefName"`
	Target string `json:"targetRefName"`
}

type GetPRsResponseUser struct {
	DisplayName string `json:"displayName"`
	Email       string `json:"uniqueName"`
}

type GetPRsResponseRepository struct {
	ID      string                `json:"id"`
	Name    string                `json:"name"`
	URL     string                `json:"url"`
	Project GetPRsResponseProject `json:"project"`
}

type GetPRsResponseProject struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ToPullRequest converts a GetPRsResponsePullRequest to a models.PullRequest
func (azPR GetPRsResponsePullRequest) ToPullRequest() models.PullRequest {
	return models.PullRequest{
		ID:             azPR.ID,
		Title:          azPR.Title,
		Description:    azPR.Description,
		RepositoryID:   azPR.Repo.ID,
		RepositoryName: azPR.Repo.Name,
		RepositoryURL:  azPR.Repo.URL,
		ProjectID:      azPR.Repo.Project.ID,
		ProjectName:    azPR.Repo.Project.Name,
		Status:         azPR.Status,
		MergeStatus:    azPR.MergeStatus,
		CreatedBy:      azPR.CreatedBy.DisplayName,
		URL:            azPR.URL,
		IsDraft:        azPR.IsDraft,
		Created:        azPR.Dates.Created,
	}
}
