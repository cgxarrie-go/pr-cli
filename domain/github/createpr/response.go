package createpr

import (
	"strconv"

	"github.com/cgxarrie-go/prq/domain/models"
)

// Response .
type Response struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"body"`
	URL         string  `json:"html_url"`
	Number		int     `json:"number"` 
	IsDraft		bool	`json:"draft"` 	

}

func (r Response) ToPullRequest(organization string) models.CreatedPullRequest {
	return models.CreatedPullRequest{
		ID:           strconv.Itoa(r.Number),
		Title:        r.Title,
		Description:  r.Description,
		URL:          r.URL,
		IsDraft:      r.IsDraft,
		Organization: organization,
		Link:         "",
	}
}
