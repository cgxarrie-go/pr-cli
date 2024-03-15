package ports

type CreateSvcPRRequest struct {
	Destination string
	Title       string
	IsDraft     bool
	PRTemplate  string
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
