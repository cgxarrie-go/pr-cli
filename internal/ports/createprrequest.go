package ports

type CreatePRRequest struct {
	Destination string
	Title       string
	IsDraft     bool
	PRTemplate  string
}
