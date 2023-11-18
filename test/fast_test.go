package test

import(
	"testing"
	"github.com/begopher/rule"
	str "github.com/begopher/rule/strings"
)


type Student struct {
	Fname string
	Lname string 
}


func Test_Map(t *testing.T) {
	nameRule := rule.All[string](
		str.Empty(),
	)
	var r rule.Rule[Student] = rule.Map[Student, string](
		nameRule,
		func(s Student) string {return s.Fname},
	)
	if ok := r.Evaluate(Student{}); !ok {
		t.Error("First name cannot be empty")
	}
	if ok := r.Evaluate(Student{Fname: "any"}); ok {
		t.Error("First name is not empty")
	}
}

func Test_Less(t *testing.T) {
	r := str.Less(3)
	if ok := r.Evaluate("AB"); !ok {
		t.Errorf("Invalid counter")
	}
	if ok := r.Evaluate("ABC"); ok {
		t.Errorf("Invalid counter")
	}
}

func Test_More(t *testing.T) {
	r := str.More(3)
	if ok := r.Evaluate("Add"); ok {
		t.Errorf("Invalid counter 1")
	}
	if ok := r.Evaluate("ABCD"); !ok {
		t.Errorf("Invalid counter 2")
	}
}


func Test_Negate(t *testing.T) {
	r := rule.Negate(str.Empty())
	if ok := r.Evaluate(""); ok {
		t.Errorf("Negate must change true to false")
	}
	if ok := r.Evaluate("any"); !ok {
		t.Errorf("Negate must change false to true")
	}
}

func Test_Is(t *testing.T) {
	is := str.Is("C", "C++")
	if ok := is.Evaluate("Python"); ok {
		t.Errorf("the value Python must return false")
	}
	if ok := is.Evaluate("C"); !ok {
		t.Errorf("the value C must return true")
	}
	if ok := is.Evaluate("C++"); !ok {
		t.Errorf("the value C++ must return true")
	}
}
