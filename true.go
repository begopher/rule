package rule

func True[T any]() Rule[T] {
	return _true[T]{}
}

type _true[T any] struct{}

func (_true[T]) Evaluate(T) bool {
	return true
}
