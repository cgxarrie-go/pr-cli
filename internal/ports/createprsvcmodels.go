package ports

type CreatePRSvcRequest struct {
	Destination string
	Title       string
	IsDraft     bool
	Template    string
	Description string
}

type CreatePRSvcResponse struct {
	ID          string
	Title       string
	Description string
	URL         string
	IsDraft     bool
	Repository  string
	Link        string
}
