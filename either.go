package rule

func Either[T any](rules ...Rule[T]) Rule[T] {
	if len(rules) == 0 {
		panic("rule.Either: rules cannot be empty")
	}
	return either[T]{rules}
}

type either[T any] struct {
	rules []Rule[T]
}

func (e either[T]) Evaluate(value T) bool {
	for _, rule := range e.rules {
		if ok := rule.Evaluate(value); ok {
			return true
		}
	}
	return false
}
