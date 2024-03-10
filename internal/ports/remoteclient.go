package ports

type RemoteClient interface {
	Do(req RemoteClientDoRequest) error
}

type RemoteClientDoRequest struct {
	Source      string
	Destination string
	Title       string
	IsDraft     bool
	Description string
}
