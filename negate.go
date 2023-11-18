package rule

func Negate[T any](rule Rule[T]) Rule[T] {
	if rule == nil {
		panic("rule.Negate: rule cannot be empty")
	}
	return negate[T]{rule}
}

type negate[T any] struct {
	rule Rule[T]
}

func (n negate[T]) Evaluate(value T) bool {
	return !n.rule.Evaluate(value)
}
