package conv

import (
	"testing"
)

func assertInflate(t *testing.T, c *Converter, input string, expected []string) {
	var actual []string
	ch := c.Inflate(input)
	for s := range ch {
		actual = append(actual, s)
	}

	if len(actual) != len(expected) {
		t.Errorf("length not match, expected=%d actual=%d",
			len(expected), len(actual))
		t.Logf("  expected=%v", expected)
		return
	}
	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("item[%d]=%v is not match with %v", i, v, expected[i])
			t.Logf("  expected=%v", expected)
			break
		}
	}
}

func TestInflate(t *testing.T) {
	c := New()
	c.Add("a", "あ", "")
	c.Add("i", "い", "")
	c.Add("u", "う", "")
	c.Add("e", "え", "")
	c.Add("o", "お", "")
	c.Add("ka", "か", "")
	c.Add("ki", "き", "")
	c.Add("ku", "く", "")
	c.Add("ke", "け", "")
	c.Add("ko", "こ", "")
	c.Add("kk", "っ", "k")

	assertInflate(t, c, "a", []string{"あ"})
	assertInflate(t, c, "ak", []string{
		"あか", "あけ", "あき", "あっ", "あこ", "あく",
	})
	assertInflate(t, c, "ik", []string{
		"いか", "いけ", "いき", "いっ", "いこ", "いく",
	})
	assertInflate(t, c, "uk", []string{
		"うか", "うけ", "うき", "うっ", "うこ", "うく",
	})
	assertInflate(t, c, "ek", []string{
		"えか", "えけ", "えき", "えっ", "えこ", "えく",
	})
	assertInflate(t, c, "ok", []string{
		"おか", "おけ", "おき", "おっ", "おこ", "おく",
	})
}
