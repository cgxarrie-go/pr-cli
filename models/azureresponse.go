package models

import "time"

type AzSvcGetPullRequestsResponse struct {
	Value []AzSvcGetPullRequestsResponsePullRequest `json:"value"`
	Count int                                       `json:"count"`
}

type AzSvcGetPullRequestsResponsePullRequest struct {
	ID          int                                    `json:"pullRequestId"`
	Title       string                                 `json:"title"`
	Description string                                 `json:"description"`
	Repo        AzSvcGetPullRequestsResponseRepository `json:"repository"`
	Status      string                                 `json:"status"`
	MergeStatus string                                 `json:"mergeStatus"`
	CreatedBy   AzSvcGetPullRequestsResponseUser       `json:"createdBy"`
	URL         string                                 `json:"url"`
	Branches    AzSvcGetPullRequestsResponseBranches
	IsDraft     bool `json:"isDraft"`
	Dates       AzSvcGetPullRequestsResponseDates
	Reviewers   []AzSvcGetPullRequestsResponseUser `json:"reviewers"`
}

type AzSvcGetPullRequestsResponseDates struct {
	Created time.Time `json:"creationDate"`
	Closed  time.Time `json:"closedDate"`
}
type AzSvcGetPullRequestsResponseBranches struct {
	Source string `json:"sourceRefName"`
	Target string `json:"targetRefName"`
}

type AzSvcGetPullRequestsResponseUser struct {
	DisplayName string `json:"displayName"`
	Email       string `json:"uniqueName"`
}

type AzSvcGetPullRequestsResponseRepository struct {
	ID      string                              `json:"id"`
	Name    string                              `json:"name"`
	URL     string                              `json:"url"`
	Project AzSvcGetPullRequestsResponseProject `json:"project"`
}

type AzSvcGetPullRequestsResponseProject struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
