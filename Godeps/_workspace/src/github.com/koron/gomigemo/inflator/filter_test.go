package inflator

import (
	"testing"
)

func TestFilter(t *testing.T) {
	p := Prefix("a", "aa", "A", "AA")
	s := Suffix("b", "bb", "B", "BB")
	j1 := Join(p, s)
	f := Filter(func(s string) bool{
		return len(s) >= 4
	})
	j2 := Join(j1, f)

	c := j2.Inflate("")
	if v, ok := <-c; !ok || v != "aabb" {
		t.Error("didn't return \"aabb\":", v, ok)
	}
	if v, ok := <-c; !ok || v != "aaBB" {
		t.Error("didn't return \"aaBB\":", v, ok)
	}
	if v, ok := <-c; !ok || v != "AAbb" {
		t.Error("didn't return \"AAbb\":", v, ok)
	}
	if v, ok := <-c; !ok || v != "AABB" {
		t.Error("didn't return \"AABB\":", v, ok)
	}
	if v, ok := <-c; ok {
		t.Error("returned unexpected:", v, ok)
	}
}
