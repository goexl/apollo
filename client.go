package apollo

import (
	"github.com/goexl/apollo/internal/builder"
	"github.com/goexl/apollo/internal/core"
)

// Client 客户端
type Client = core.Client

func New() *builder.Client {
	return builder.NewClient()
}
