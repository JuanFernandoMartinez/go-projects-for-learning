package greetings

import (
	"regexp"
	"testing"
)

func TestHeloName(t *testing.T) {
	name := "Juan"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
