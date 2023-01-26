package models

import (
	"time"
)

// PullRequest is the abstraction of a Pull Request from any provider
type PullRequest struct {
	ID             string
	Title          string
	Description    string
	RepositoryID   string
	RepositoryName string
	RepositoryURL  string
	ProjectID      string
	ProjectName    string
	Status         string
	MergeStatus    string
	CreatedBy      string
	URL            string
	IsDraft        bool
	Created        time.Time
}
