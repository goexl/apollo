package builder

import (
	"github.com/goexl/apollo/internal/param"
)

type Base[T any] struct {
	from   *T
	params *param.Base
}

func NewBase[T any](from *T, params *param.Base) *Base[T] {
	return &Base[T]{
		from:   from,
		params: params,
	}
}

func (b *Base[T]) Cluster(cluster string) (t *T) {
	b.params.Cluster = cluster
	t = b.from

	return
}

func (b *Base[T]) Namespace(namespace string) (t *T) {
	b.params.Namespaces = append(b.params.Namespaces, namespace)
	t = b.from

	return
}

func (b *Base[T]) Namespaces(namespace string, namespaces ...string) (t *T) {
	b.params.Namespaces = append(b.params.Namespaces, append([]string{namespace}, namespaces...)...)
	t = b.from

	return
}
