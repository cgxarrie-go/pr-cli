package models

import (
	"fmt"
	"time"
)

// PullRequest is the abstraction of a Pull Request from any provider
type PullRequest struct {
	ID          string
	Title       string
	Description string
	Repository  Hierarchy
	Project     Hierarchy
	Status      string
	MergeStatus string
	CreatedBy   string
	URL         string
	IsDraft     bool
	Created     time.Time
}

type CreatedPullRequest struct {
	ID           string
	Title        string
	Description  string
	Repository   Hierarchy
	Project      Hierarchy
	URL          string
	IsDraft      bool
	Organization string
}

// Hierarchy of a PR
type Hierarchy struct {
	ID   string
	Name string
	URL  string
}

// ShortenedTitle returns title shortened to maxlength...
func (p PullRequest) ShortenedTitle(maxLength int) string {

	if len(p.Title) <= maxLength {
		return p.Title
	}

	title := fmt.Sprintf("%s...", p.Title[0:maxLength-3])
	return title
}
