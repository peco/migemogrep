package inflator

import (
	"testing"
)

func TestDispatch(t *testing.T) {
	e := Echo()
	p := Prefix("foo1-", "foo2-")
	s := Suffix("-bar1", "-bar2")
	d := Dispatch(e, p, s)

	c1 := d.Inflate("qux")
	if v, ok := <-c1; !ok || v != "qux" {
		t.Error("didn't return \"qux\":", v, ok)
	}
	if v, ok := <-c1; !ok || v != "foo1-qux" {
		t.Error("didn't return \"foo1-qux\":", v, ok)
	}
	if v, ok := <-c1; !ok || v != "foo2-qux" {
		t.Error("didn't return \"foo2-qux\":", v, ok)
	}
	if v, ok := <-c1; !ok || v != "qux-bar1" {
		t.Error("didn't return \"qux-bar1\":", v, ok)
	}
	if v, ok := <-c1; !ok || v != "qux-bar2" {
		t.Error("didn't return \"qux-bar2\":", v, ok)
	}
	if v, ok := <-c1; ok {
		t.Error("returned unexpected:", v, ok)
	}

	c2 := d.Inflate("baz")
	if v, ok := <-c2; !ok || v != "baz" {
		t.Error("didn't return \"baz\":", v, ok)
	}
	if v, ok := <-c2; !ok || v != "foo1-baz" {
		t.Error("didn't return \"foo1-baz\":", v, ok)
	}
	if v, ok := <-c2; !ok || v != "foo2-baz" {
		t.Error("didn't return \"foo2-baz\":", v, ok)
	}
	if v, ok := <-c2; !ok || v != "baz-bar1" {
		t.Error("didn't return \"baz-bar1\":", v, ok)
	}
	if v, ok := <-c2; !ok || v != "baz-bar2" {
		t.Error("didn't return \"baz-bar2\":", v, ok)
	}
	if v, ok := <-c2; ok {
		t.Error("returned unexpected:", v, ok)
	}
}
