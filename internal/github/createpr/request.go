package createpr

import "github.com/cgxarrie-go/prq/internal/ports"

type Request struct {
	destination string
	title       string
}

// Destination implements ports.CreatePRRequest.
func (r Request) Destination() string {
	return r.destination
}

// Title implements ports.CreatePRRequest.
func (r Request) Title() string {
	return r.title
}

func NewRequest(destination, title string) ports.CreatePRRequest {
	return Request{
		destination: destination,
		title:       title,
	}
}
