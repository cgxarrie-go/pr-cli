package models

import (
	"fmt"
	"time"
)

// PullRequest is the abstraction of a Pull Request from any provider
type PullRequest struct {
	Origin 		 string
	Organization string
	ID           string
	Title        string
	Description  string
	Repository   Hierarchy
	Project      Hierarchy
	Status       string
	MergeStatus  string
	CreatedBy    string
	URL          string
	IsDraft      bool
	Created      time.Time
	Link 		 string
}

type CreatedPullRequest struct {
	ID           string
	Title        string
	Description  string
	URL          string
	IsDraft      bool
	Organization string
	Link 		 string
}

// Hierarchy of a PR
type Hierarchy struct {
	ID   string
	Name string
	URL  string
}

// ShortenedTitle returns title shortened to maxlength...
func (p PullRequest) ShortenedTitle(maxLength int) string {

	draftText := ""
	if p.IsDraft {
		draftText = " (Draft)"
	}

	pritntable := fmt.Sprintf("%s%s", p.Title, draftText)

	if len(pritntable) <= maxLength {
		return pritntable
	}



	shortenLenght := maxLength - 3 - len(draftText)

	title := fmt.Sprintf("%s...", pritntable[0:shortenLenght])
	return title
}
