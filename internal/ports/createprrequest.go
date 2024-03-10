package ports

type CreatePRRequest interface {
	Destination() string
	Title() string
	IsDraft() bool
}
