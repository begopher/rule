package constraint

import(
	"github.com/begopher/rule"
)

func New[T any](rule rule.Rule[T], err error) rule.Constraint[T] {
	if rule == nil {
		panic("constraint.New: rule cannot be nil")
	}
	if err == nil {
		panic("constraint.New: err cannot be nil")
	}
	return _constraint[T]{rule, err}
}

type _constraint[T any] struct{
	rule rule.Rule[T]
	err error
}

func (c _constraint[T]) Evaluate(value T) error {
	if c.rule.Evaluate(value) {
		return c.err
	}
	return nil
}
