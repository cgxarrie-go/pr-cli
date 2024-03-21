package remotetype

type RemoteType uint8

const (
	NotSet RemoteType = 0
	Azure  RemoteType = 1
	Github RemoteType = 2
)

var names = map[RemoteType]string{
	NotSet: `not-set`,
	Azure:  `dev.azure.com`,
	Github: `github.com`,
}

// Name return the name of the RemoteType
func (t RemoteType) Name() string {
	return names[t]
}

// FromName returns a names matching the provided name.
func FromName(name string) RemoteType {
	for id, n := range names {
		if n == name {
			return id
		}
	}

	return NotSet
}
