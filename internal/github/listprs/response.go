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
	ID          int        `json:"id"`
	URL         string     `json:"html_url"`
	Number      int 	   `json:"number"`
	Title       string     `json:"title"`
	Body       string      `json:"body"`
	Status      string     `json:"sate"`
	MergeStatus string     `json:"mergeStatus"`
	User   ResponseUser    `json:"user"`
	IsDraft     bool       `json:"draft"`
	Created     time.Time  `json:"created_at"`
	Closed      time.Time  `json:"closed_at"`
}

// ResponseUser Pull request response user
type ResponseUser struct {
	Login string `json:"login"`
}



// ToPullRequest converts a GetPRsResponsePullRequest to a models.PullRequest
func (ghPR ResponsePullRequest) ToPullRequest(organization string) models.PullRequest {
	return models.PullRequest{
		Orgenization: organization,
		ID:           strconv.Itoa(ghPR.Number),
		Title:        ghPR.Title,
		Description:  ghPR.Body,
		Status:      ghPR.Status,
		CreatedBy:   ghPR.User.Login,
		URL:         ghPR.URL,
		IsDraft:     ghPR.IsDraft,
		Created:     ghPR.Created,
	}
}