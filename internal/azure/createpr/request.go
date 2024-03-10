package createpr

import "github.com/cgxarrie-go/prq/internal/ports"

type Request struct {
	destination string
	title       string
	isDraft     bool
}

// Destination implements ports.CreatePRRequest.
func (r Request) Destination() string {
	return r.destination
}

// Title implements ports.CreatePRRequest.
func (r Request) Title() string {
	return r.title
}

// IsDraft implements ports.CreatePRRequest.
func (r Request) IsDraft() bool {
	return r.isDraft
}

func NewRequest(destination, title string, isDraft bool) ports.CreatePRRequest {
	return Request{
		destination: destination,
		title:       title,
		isDraft:     isDraft,
	}
}
