package ports

import (
	"github.com/cgxarrie-go/prq/internal/utils"
)

type ListPRRequest interface {
	Origins() utils.Origins
}