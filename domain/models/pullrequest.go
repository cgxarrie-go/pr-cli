package models

import "time"

type PullRequest struct {
	ID             int
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
