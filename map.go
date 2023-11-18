package rule

func Map[F, T any](rule Rule[T], mapping func(F) T) Rule[F] {
	if rule == nil {
		panic("str.Map: rule cannot be nil")
	}
	if mapping == nil {
		panic("str.Map: mapping cannot be nil")
	}
	return _map[F,T]{rule, mapping}
}

type _map[F, T any] struct {
	rule Rule[T]
	mapping   func(F) T
}

func (m _map[F, T]) Evaluate(from F) bool {
	to := m.mapping(from)
	return m.rule.Evaluate(to)
}
