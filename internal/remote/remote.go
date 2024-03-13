package remote

import (
	"github.com/cgxarrie-go/prq/internal/models"
	"github.com/cgxarrie-go/prq/internal/remotetype"
)

type Remote struct {
	remote              string
	remoteType          remotetype.RemoteType
	defaultTargetBranch models.Branch
}

func NewRemote(r string, t remotetype.RemoteType,
	defTgtBranch models.Branch) Remote {

	return Remote{
		remote:              r,
		remoteType:          t,
		defaultTargetBranch: defTgtBranch,
	}

}

func (r Remote) DefaultTargetBranch() models.Branch {
	return r.defaultTargetBranch
}

func (r Remote) String() string {
	return r.remote
}

type Remotes map[Remote]any

func (rs *Remotes) Append(r Remote) {

	if _, ok := (*rs)[r]; ok {
		return
	}

	(*rs)[r] = struct{}{}

}
