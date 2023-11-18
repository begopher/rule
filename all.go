package rule

func All[T any](rules ...Rule[T]) Rule[T] {
	if len(rules) == 0 {
		panic("rule.All: rules cannot be empty")
	}
	return all[T]{rules}
}

type all[T any] struct {
	rules []Rule[T]
}

func (a all[T]) Evaluate(value T) bool {
	for _, rule := range a.rules {
		if ok := rule.Evaluate(value); !ok {
			return false
		}
	}
	return true
}


