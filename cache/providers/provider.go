package providers

import "fmt"

type Provider int

const (
	NotSet Provider = 0
	Azure  Provider = 1
)

var providerMapName = map[Provider]string{
	NotSet: `not-set`,
	Azure:  `az`,
}

// CommandName return the name of the Provider
func (t Provider) CommandName() string {
	return providerMapName[t]
}

// FromName returns a providerMapName matching the provided name.
func FromName(name string) (Provider, error) {
	for id, n := range providerMapName {
		if n == name {
			return id, nil
		}
	}

	return NotSet, fmt.Errorf("")
}
