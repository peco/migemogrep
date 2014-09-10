package readutil

import (
	"io"
	"testing"
)

func assertReadRune(t *testing.T, exp rune, r *StackableRuneReader) {
	ch, _, err := r.ReadRune()
	if err != nil {
		t.Error(err)
	} else if ch != exp {
		t.Errorf("rune mismatch, expected=%c actual=%c", exp, ch)
	}
}

func assertEOF(t *testing.T, r *StackableRuneReader) {
	_, _, err := r.ReadRune()
	if err == nil {
		t.Error("expected io.EOF but no error actually")
	} else if err != io.EOF {
		t.Error("expected io.EOF but actual: ", err)
	}
}

func TestStackedReader(t *testing.T) {
	r := NewStackabeRuneReader()
	r.PushFront("abc")
	assertReadRune(t, 'a', r)
	assertReadRune(t, 'b', r)
	assertReadRune(t, 'c', r)
	assertEOF(t, r)
}

func TestStackedReaderMultiple(t *testing.T) {
	r := NewStackabeRuneReader()
	r.PushFront("foo")
	r.PushFront("bar")
	assertReadRune(t, 'b', r)
	assertReadRune(t, 'a', r)
	assertReadRune(t, 'r', r)
	assertReadRune(t, 'f', r)
	assertReadRune(t, 'o', r)
	assertReadRune(t, 'o', r)
	assertEOF(t, r)
}

func TestStackedReaderMultipleSuspend1(t *testing.T) {
	r := NewStackabeRuneReader()
	r.PushFront("foo")
	assertReadRune(t, 'f', r)
	r.PushFront("bar")
	assertReadRune(t, 'b', r)
	assertReadRune(t, 'a', r)
	assertReadRune(t, 'r', r)
	assertReadRune(t, 'o', r)
	assertReadRune(t, 'o', r)
	assertEOF(t, r)
}

func TestStackedReaderMultipleSuspend2(t *testing.T) {
	r := NewStackabeRuneReader()
	r.PushFront("foo")
	assertReadRune(t, 'f', r)
	assertReadRune(t, 'o', r)
	r.PushFront("bar")
	assertReadRune(t, 'b', r)
	assertReadRune(t, 'a', r)
	assertReadRune(t, 'r', r)
	assertReadRune(t, 'o', r)
	assertEOF(t, r)
}

func TestStackedReaderMultipleSuspend3(t *testing.T) {
	r := NewStackabeRuneReader()
	r.PushFront("foo")
	assertReadRune(t, 'f', r)
	assertReadRune(t, 'o', r)
	assertReadRune(t, 'o', r)
	r.PushFront("bar")
	assertReadRune(t, 'b', r)
	assertReadRune(t, 'a', r)
	assertReadRune(t, 'r', r)
	assertEOF(t, r)
}
