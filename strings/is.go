package strings

import(
	"github.com/begopher/rule"
)

func Is(many ...string) rule.Rule[string] {
	if len(many) == 0 {
		panic("strings.Is: many cannot be empty")
	}
	return is{many}
}

type is struct {
	many []string
}

func (i is) Evaluate(value string) bool {
	for _, val := range i.many {
		if val == value {
			return true
		}
	}
	return false
}
