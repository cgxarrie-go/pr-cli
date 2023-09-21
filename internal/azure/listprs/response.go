package listprs

import (
	"strconv"
	"time"

	"github.com/cgxarrie-go/prq/internal/models"
)

// Response response from GetPRS fro Azure Service
type Response struct {
	Value []ResponsePullRequest `json:"value"`
	Count int                         `json:"count"`
}

// ResponsePullRequest Pull request response item
type ResponsePullRequest struct {
	ID          int                      `json:"pullRequestId"`
	Title       string                   `json:"title"`
	Description string                   `json:"description"`
	Repo        ResponseRepository `json:"repository"`
	Status      string                   `json:"status"`
	MergeStatus string                   `json:"mergeStatus"`
	CreatedBy   ResponseUser       `json:"createdBy"`
	URL         string                   `json:"url"`
	Branches    ResponseBranches
	IsDraft     bool                 `json:"isDraft"`
	Reviewers   []ResponseUser `json:"reviewers"`
	Created     time.Time            `json:"creationDate"`
	Closed      time.Time            `json:"closedDate"`
}

// ResponseBranches pull request response branches
type ResponseBranches struct {
	Source string `json:"sourceRefName"`
	Target string `json:"targetRefName"`
}

// ResponseUser Pull request response user
type ResponseUser struct {
	DisplayName string `json:"displayName"`
	Email       string `json:"uniqueName"`
}

// ResponseRepository pull request response repository
type ResponseRepository struct {
	ID      string                `json:"id"`
	Name    string                `json:"name"`
	URL     string                `json:"url"`
	Project ResponseProject `json:"project"`
}

// ResponseProject pull request response project
type ResponseProject struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// ToPullRequest converts a GetPRsResponsePullRequest to a models.PullRequest
func (azPR ResponsePullRequest) ToPullRequest() models.PullRequest {
	return models.PullRequest{
		ID:           strconv.Itoa(azPR.ID),
		Title:        azPR.Title,
		Description:  azPR.Description,
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