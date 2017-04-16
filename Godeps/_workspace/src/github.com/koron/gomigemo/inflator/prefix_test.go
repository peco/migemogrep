package inflator

import (
	"testing"
)

func TestPrefix(t *testing.T) {
	p := Prefix("foo", "bar", "baz")
	c := p.Inflate("-qux")
	if "foo-qux" != <-c {
		t.Error("Prefix didn't return \"foo-qux\"")
	}
	if "bar-qux" != <-c {
		t.Error("Prefix didn't return \"bar-qux\"")
	}
	if "baz-qux" != <-c {
		t.Error("Prefix didn't return \"baz-qux\"")
	}
	if _, ok := <-c; ok {
		t.Error("Prefix returned unexpected")
	}
}
