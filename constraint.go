package rule

type Constraint[T any] interface {
	Evaluate(T) error
}
