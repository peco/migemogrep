package inflator

import (
	"testing"
)

func TestSuffix(t *testing.T) {
	s := Suffix("foo", "bar", "baz")
	c := s.Inflate("qux-")
	if "qux-foo" != <-c {
		t.Error("Suffix didn't return \"qux-foo\"")
	}
	if "qux-bar" != <-c {
		t.Error("Suffix didn't return \"qux-bar\"")
	}
	if "qux-baz" != <-c {
		t.Error("Suffix didn't return \"qux-baz\"")
	}
	if _, ok := <-c; ok {
		t.Error("Suffix returned unexpected")
	}
}
