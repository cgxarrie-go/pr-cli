package status

import "fmt"

type Status int

const (
	Active    Status = 1
	Abandoned Status = 2
	Closed    Status = 3
)

var enumMapName = map[Status]string{
	Active:    `active`,
	Abandoned: `abandoned`,
	Closed:    `cancelled`,
}

// Name return the name of the Status
func (t Status) Name() string {
	return enumMapName[t]
}

// FromName returns a enumMapName matching the provided name.
func FromName(name string) (Status, error) {
	for id, n := range enumMapName {
		if n == name {
			return id, nil
		}
	}

	return Active, fmt.Errorf("not a valid status %s", name)
}
