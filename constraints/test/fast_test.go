package test

import(
	"fmt"
	"testing"
	"github.com/begopher/rule"
	"github.com/begopher/rule/strings"
	"github.com/begopher/rule/constraint"
	"github.com/begopher/rule/constraints"
)


func Test_constraints(t *testing.T) {
	c := constraints.New[string]()
	if err := c.Evaluate(""); err != nil {
		t.Errorf("No constraint must return nil, got (%v)", err)
	}
	violationErr := fmt.Errorf("violation-error")
	duplicatedErr := fmt.Errorf("duplicated-error")
	pc := constraints.New(
		constraint.New(
			rule.Either(strings.Empty(), strings.More(4)),
			violationErr,
		),
		constraint.New(
			strings.Is("C", "C++"),
			duplicatedErr,
		),
	)
	if err := pc.Evaluate(""); err != violationErr {
		t.Errorf("Expected error is (%v) got (%v)", violationErr, err)
	}
	if err := pc.Evaluate("more than 4 characters"); err != violationErr {
		t.Errorf("Expected error is (%v) got (%v)", violationErr, err)
	}
	if err := pc.Evaluate("C"); err != duplicatedErr {
		t.Errorf("Expected error is (%v) got (%v)", duplicatedErr, err)
	}
	if err := pc.Evaluate("C++"); err != duplicatedErr {
		t.Errorf("Expected error is (%v) got (%v)", duplicatedErr, err)
	}
	if err := pc.Evaluate("any"); err != nil {
		t.Errorf("Expected error is (nil) got (%v)",  err)
	}
}
