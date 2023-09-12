package utils

import "strings"

type Origin string
type Origins []Origin

func (o Origin) IsGithub() bool {
	return strings.Contains(string(o), "github.com")
}

func (o Origin) IsAzure() bool {
	return strings.Contains(string(o), "dev.azure.com")
}

func (o Origins) Append(origin Origin) Origins {
	exists := false
	for _, o := range o {
		if o == origin {
			exists = true
			break
		}
	}
	// append origin if it doesn't exist
	if !exists {
		return append(o, origin)
	}
	return o
}
