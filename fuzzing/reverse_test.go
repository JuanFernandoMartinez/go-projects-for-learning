package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}

	for _, tc := range testCases {
		rev, revErr := Reverse(tc.in)
		if revErr != nil {
			return
		}
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", tc.in, tc.want)
		}
	}

}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev, revErr := Reverse(orig)
		doubleRev, dobRevErr := Reverse(rev)
		if revErr != nil {
			return
		}

		if dobRevErr != nil {
			return
		}

		if orig != doubleRev {
			t.Errorf("Before: %q, Afer: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
