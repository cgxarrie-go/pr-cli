package azure

import (
	"github.com/cgxarrie/pr-go/domain/models"
	"github.com/cgxarrie/pr-go/domain/ports"
)

type azureSvc struct {
	conpanyName string
	pat         string
}

// GetPRs implements ports.ProviderService
func (azureSvc) GetPRs() (prs []models.PullRequest, err error) {
	panic("unimplemented")
}

// NewAzureService return new instnce of azure service
func NewAzureService(companyName string, pat string) ports.ProviderService {
	return azureSvc{
		conpanyName: companyName,
		pat:         pat,
	}
}
