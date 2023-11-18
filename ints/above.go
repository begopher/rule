package ints

import(
	"github.com/begopher/rule"
)

func Above(number int) rule.Rule[int] {
	return above{number}
}

type above struct {
	number int
}

func (a above) Evaluate(value int) bool {
	if value > a.number {
		return true
	}
	return false
}
