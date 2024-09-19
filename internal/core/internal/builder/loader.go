package builder

import (
	"github.com/goexl/apollo/internal/core/internal/param"
)

type Loader struct {
	params *param.Loader
}

func NewLoader(params *param.Client) *Loader {
	return &Loader{
		params: param.NewLoader(params),
	}
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
