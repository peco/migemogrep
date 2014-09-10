package inflator

import (
	"testing"
)

func TestEcho(t *testing.T) {
	e := Echo()

	c1 := e.Inflate("foo")
	if "foo" != <-c1 {
		t.Error("Echo didn't return \"foo\"")
	}
	if _, ok := <-c1; ok {
		t.Error("Echo returned others of \"foo\"")
	}

	c2 := e.Inflate("bar")
	if "bar" != <-c2 {
		t.Error("Echo didn't return \"bar\"")
	}
	if _, ok := <-c2; ok {
		t.Error("Echo returned others of \"bar\"")
	}
}
