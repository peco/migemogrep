package conv

import (
	"testing"
)

func assertConvert(t *testing.T, c *Converter, expected, input string) {
	converted, err := c.Convert(input)
	if err != nil {
		t.Error("Convert failed:", err)
		t.Logf("  input=%s expected=%s", input, expected)
	}
	if converted != expected {
		t.Error("Convert returns unexpected:", converted)
		t.Logf("  input=%s expected=%s", input, expected)
	}
}

func TestEmpty(t *testing.T) {
	c := New()
	assertConvert(t, c, "", "")
	assertConvert(t, c, "foo", "foo")
	assertConvert(t, c, "bar", "bar")
}

func TestSimple(t *testing.T) {
	c := New()
	c.Add("a", "A", "")
	c.Add("b", "B", "")
	assertConvert(t, c, "A", "a")
	assertConvert(t, c, "B", "b")
	assertConvert(t, c, "c", "c")
	assertConvert(t, c, "AAAAABBBBBccccc", "aaaaabbbbbccccc")
}

func TestTiny(t *testing.T) {
	c := New()
	c.Add("aa", "A", "a")
	c.Add("ab", "B", "")
	assertConvert(t, c, "B", "ab")
	assertConvert(t, c, "Aa", "aa")
	assertConvert(t, c, "AB", "aab")
	assertConvert(t, c, "Bc", "abc")
	assertConvert(t, c, "Ba", "aba")
}

func TestHira(t *testing.T) {
	c := New()
	c.Add("a", "あ", "")
	c.Add("i", "い", "")
	assertConvert(t, c, "あい", "ai")
}

func TestHiraRemain(t *testing.T) {
	c := New()
	c.Add("a", "あ", "")
	c.Add("ka", "か", "")
	c.Add("ki", "き", "")
	assertConvert(t, c, "あk", "ak")
}
