package status

import "fmt"

type Status int

const (
	Active    Status = 1
	Abandoned Status = 2
	Cancelled Status = 3
)

var statusMapName = map[Status]string{
	Active:    `active`,
	Abandoned: `abandoned`,
	Cancelled: `cancelled`,
}

// Name return the name of the Status
func (t Status) Name() string {
	return statusMapName[t]
}

// FromName returns a statusMapName matching the provided name.
func FromName(name string) (Status, error) {
	for id, n := range statusMapName {
		if n == name {
			return id, nil
		}
	}

	return Active, fmt.Errorf("not a valid status %s", name)
}
