package branch

import (
	"fmt"
	"strings"
)

const prefix = "refs/heads/"

type Branch struct {
	name string
}

func NewBranch(name string) Branch {
	return Branch{
		name: strings.TrimPrefix(name, prefix),
	}
}

func (b Branch) Name() string {
	return b.name
}

func (b Branch) FullName() string {
	return fmt.Sprintf("%s%s", prefix, b.name)
}
