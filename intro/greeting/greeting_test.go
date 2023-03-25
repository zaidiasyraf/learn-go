package greeting

import (
	"regexp"
	"testing"
)

func TestHelloNamet(t *testing.T) {
	name := "Ken"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Ken")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Ken") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloNameEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
