package createpr

import (
	"strconv"

	"github.com/cgxarrie-go/prq/internal/models"
)

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


// Response .
type Response struct {
	ID          int                      `json:"pullRequestId"`
	Title       string                   `json:"title"`
	Description string                   `json:"description"`
	Repo        ResponseRepository `json:"repository"`
	URL         string                   `json:"url"`
	IsDraft     bool                     `json:"isDraft"`
}

func (azPR Response) ToPullRequest(organization string) models.CreatedPullRequest {
	return models.CreatedPullRequest{
		ID:          strconv.Itoa(azPR.ID),
		Title:       azPR.Title,
		Description: azPR.Description,
		URL:          azPR.URL,
		IsDraft:      azPR.IsDraft,
		Organization: organization,
	}
}
