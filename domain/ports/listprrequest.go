package ports

import (
	"github.com/cgxarrie-go/prq/utils"
)

type ListPRRequest interface {
	Origins() utils.Origins
	Status()  PRStatus
}