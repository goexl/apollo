package builder

import (
	"github.com/goexl/apollo/internal/param"
)

type base[T any] struct {
	from   *T
	params *param.Base
}

func newBase[T any](from *T, params *param.Base) *base[T] {
	return &base[T]{
		from:   from,
		params: params,
	}
}

func (b *base[T]) Cluster(cluster string) (t *T) {
	b.params.Cluster = cluster
	t = b.from

	return
}

func (b *base[T]) Namespace(namespace string) (t *T) {
	b.params.Namespaces = append(b.params.Namespaces, namespace)
	t = b.from

	return
}

func (b *base[T]) Namespaces(namespace string, namespaces ...string) (t *T) {
	b.params.Namespaces = append(b.params.Namespaces, append([]string{namespace}, namespaces...)...)
	t = b.from

	return
}
