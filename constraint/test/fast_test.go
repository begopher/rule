package test

import (
	"fmt"
	"testing"
	"github.com/begopher/rule/constraint"
)

type Student struct {
	Name string
}


func Test_Map(t *testing.T) {
	name := "ali"
	// test 1
	cons := constraint.Map[Student, string](
		func(s Student) string {return s.Name},
		playing{name},
	)
	if err := cons.Evaluate(Student{name}); err != nil {
		t.Error(err.Error())
	}

	// test 2
	cons = constraint.Map[Student, string](
		func(s Student) string {return s.Name},
		playing{name+"++"},
	)
	if err := cons.Evaluate(Student{name}); err == nil {
		t.Error(err.Error())
	}
	
}

type playing struct {
	name string
}

func (p playing) Evaluate(name string) error {
	if p.name == name {
		return nil
	}
	return fmt.Errorf("expected name (%v) got (%v)", p.name, name)
}
