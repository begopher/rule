package ints

import(
	"github.com/begopher/rule"
)

func Below(number int) rule.Rule[int] {
	return below{number}
}

type below struct {
	number int
}

func (b below) Evaluate(value int) bool {
	if value < b.number {
		return true
	}
	return false
}
