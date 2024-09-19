package param

import (
	"github.com/goexl/apollo/internal/param"
)

type Loader struct {
	*param.Client

	Key          string
	Label        string
	Notification uint64
	Ip           string
}

func NewLoader(client *param.Client) *Loader {
	return &Loader{
		Client: client,
	}
}
