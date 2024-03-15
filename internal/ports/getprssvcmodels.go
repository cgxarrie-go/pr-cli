package ports

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
	IsDraft     bool
	Link        string
}
