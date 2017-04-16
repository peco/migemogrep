package inflator

import (
	"testing"
)

func TestJoin(t *testing.T) {
	p := Prefix("foo-")
	s := Suffix("-bar")
	j := Join(p, s)
	c := j.Inflate("qux")

	if v, ok := <-c; !ok || v != "foo-qux-bar" {
		t.Error("didn't return \"foo-qux-bar\":", v, ok)
	}
	if v, ok := <-c; ok {
		t.Error("returned unexpected:", v, ok)
	}
}

func TestJoinMulti(t *testing.T) {
	p := Prefix("foo1-", "foo2-", "foo3-")
	s := Suffix("-bar1", "-bar2")
	c := Join(p, s).Inflate("qux")

	if v, ok := <-c; !ok || v != "foo1-qux-bar1" {
		t.Error("didn't return \"foo1-qux-bar1\":", v, ok)
	}
	if v, ok := <-c; !ok || v != "foo1-qux-bar2" {
		t.Error("didn't return \"foo1-qux-bar2\":", v, ok)
	}
	if v, ok := <-c; !ok || v != "foo2-qux-bar1" {
		t.Error("didn't return \"foo2-qux-bar1\":", v, ok)
	}
	if v, ok := <-c; !ok || v != "foo2-qux-bar2" {
		t.Error("didn't return \"foo2-qux-bar2\":", v, ok)
	}
	if v, ok := <-c; !ok || v != "foo3-qux-bar1" {
		t.Error("didn't return \"foo3-qux-bar1\":", v, ok)
	}
	if v, ok := <-c; !ok || v != "foo3-qux-bar2" {
		t.Error("didn't return \"foo3-qux-bar2\":", v, ok)
	}
	if v, ok := <-c; ok {
		t.Error("returned unexpected:", v, ok)
	}
}
