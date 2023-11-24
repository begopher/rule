package constraints

import(
	"github.com/begopher/rule"
)

func New[T any](consts ...rule.Constraint[T]) rule.Constraint[T] {
	for _, cons := range consts {
		if cons == nil {
			panic("constraints.New: consts cannot have nil value")
		}
	}
	return constraints[T]{consts}
}

type constraints[T any] struct {
	consts []rule.Constraint[T]
}

func (c constraints[T]) Evaluate(value T) error {
	for _, cons := range c.consts {
		if err := cons.Evaluate(value); err != nil {
			return err
		}
	}
	return nil
}

