package core

import (
	"github.com/goexl/apollo/internal/core/internal/core/internal"
	"github.com/goexl/apollo/internal/core/internal/param"
)

type Loader struct {
	params *param.Loader
	client *internal.Client
}

func NewLoader(params *param.Loader) *Loader {
	return &Loader{
		params: params,
		client: internal.NewClient(params),
	}
}

func (g *Loader) Load(target any) error {
	return g.client.Load(target, g.params)
}
