package ports

// PRReader Contract for all services reading Pull requests from providers
type PRReader interface {
	Run() (resp GetPRsSvcResponse)
}

type PRCreator interface {
	Run(req CreatePRSvcRequest) (pr CreatePRSvcResponse, err error)
}
