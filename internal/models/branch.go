package models

type Branch struct {
	name     string
	fullName string
}

func NewBranch(name, fullName string) Branch {
	return Branch{
		name:     name,
		fullName: fullName,
	}
}

func (b Branch) Name() string {
	return b.name
}

func (b Branch) FullName() string {
	return b.fullName
}
