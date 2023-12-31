package strings

import(
	"github.com/begopher/rule"
	"unicode/utf8"
)

func More(runeNum int) rule.Rule[string] {
	return more{runeNum}
}

type more struct {
	number int 
}

func (l more) Evaluate(value string) bool {
	n := utf8.RuneCountInString(value)
	if n > l.number {
		return true
	}
	return false
}
