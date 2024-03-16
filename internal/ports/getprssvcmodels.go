package ports

import "time"

type GetPRsSvcResponse struct {
	Remote       string
	Count        int
	PullRequests []GetPRsSvcResponseItem
	Error        error
}

type GetPRsSvcResponseItem struct {
	ID          string
	Title       string
	Description string
	Created     time.Time // TODO map from client response:
	CreatedBy   string    // TODO map from client response:
	IsDraft     bool
	Link        string
}
