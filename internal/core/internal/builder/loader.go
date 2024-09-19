package builder

import (
	"context"

	"github.com/goexl/apollo/internal/core/internal/core"
	"github.com/goexl/apollo/internal/core/internal/param"
	"github.com/goexl/apollo/internal/internal/builder"
)

type loader = builder.Base[Loader]

type Loader struct {
	*loader

	params *param.Loader
}

func NewLoader(params *param.Client) (loader *Loader) {
	loader = new(Loader)
	loader.params = param.NewLoader(params)

	loader.loader = builder.NewBase(loader, loader.params.Base)

	return
}

func (l *Loader) Context(ctx context.Context) (loader *Loader) {
	l.params.Context = ctx
	loader = l

	return
}

func (l *Loader) Key(key string) (loader *Loader) {
	l.params.Key = key
	loader = l

	return
}

func (l *Loader) Notification(notification uint64) (loader *Loader) {
	l.params.Notification = notification
	loader = l

	return
}

func (l *Loader) Label(label string) (loader *Loader) {
	l.params.Label = label
	loader = l

	return
}

func (l *Loader) Ip(ip string) (loader *Loader) {
	l.params.Ip = ip
	loader = l

	return
}

func (l *Loader) Build() *core.Loader {
	return core.NewLoader(l.params)
}
