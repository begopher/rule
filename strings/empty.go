package strings

import(
	"github.com/begopher/rule"
)

func Empty() rule.Rule[string] {
	return notEmpty{}
}

type notEmpty struct {}

func (n notEmpty) Evaluate(value string) bool {
	if value == "" {
		return true 
	}
	return false
}
