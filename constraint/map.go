package constraint

import(
	"github.com/begopher/rule"
)

func Map[F, T any](mapping func(F) T, cons rule.Constraint[T]) rule.Constraint[F] {
	if cons == nil {
		panic("constraint.Map: cons cannot be nil")
	}
	if mapping == nil {
		panic("constraint.Map: mapping cannot be nil")
	}
	return _map[F,T]{cons, mapping}
}

type _map[F, T any] struct {
	cons      rule.Constraint[T]
	mapping   func(F) T
}

func (m _map[F, T]) Evaluate(from F) error{
	to := m.mapping(from)
	return m.cons.Evaluate(to)
}
