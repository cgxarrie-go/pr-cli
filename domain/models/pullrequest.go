package models

import (
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

// Hierarchy of a PR
type Hierarchy struct {
	ID   string
	Name string
	URL  string
}
