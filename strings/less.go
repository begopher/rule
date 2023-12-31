package strings

import(
	"github.com/begopher/rule"
	"unicode/utf8"
)

func Less(runeNum int) rule.Rule[string] {
	return less{runeNum}
}

type less struct {
	number int 
}

func (l less) Evaluate(value string) bool {
	n := utf8.RuneCountInString(value)
	if n < l.number {
		return true
	}
	return false
}
