package errorf_test

import (
	"testing"

	"github.com/cyy0523xc/go-utils/errorf"
)

func TestNewEqual(t *testing.T) {
	// Different allocations should not be equal.
	if errorf.New("abc", "code") == errorf.New("abc", "code") {
		t.Errorf(`New("abc", "code") == New("abc", "code")`)
	}
	if errorf.New("abc", "code") == errorf.New("xyz", "code") {
		t.Errorf(`New("abc", "code") == New("xyz", "code")`)
	}

	// Same allocation should be equal to itself (not crash).
	err := errorf.New("jkl", "code")
	if err != err {
		t.Errorf(`err != err`)
	}
}

func TestErrorMethod(t *testing.T) {
	err := errorf.New(100, "abc")
	if err.Error() != "100abc" {
		t.Errorf(`New(100, "code").Error() = %q, want %q`, err.Error(), "100abc")
	}

	err = errorf.Newf("%s: %s", "abc", "code")
	if err.Error() != "abc: code" {
		t.Errorf(`New("%s: %s", "abc", "code").Error() = %q, want %q`, "abc", "code", err.Error(), "abc: code")
	}
}
