package conv

import (
	"testing"
)

func TestLoad(t *testing.T) {
	c := New()
	count, err := c.LoadFile("./load_test0.txt")
	if err != nil {
		t.Fatal("failed to load load_test0.txt", err)
	} else if count != 18 {
		t.Fatal("load_test0.txt has unexpected entries:", count)
	}

	assertConvert(t, c, "あかん", "akann")
	assertConvert(t, c, "あk", "ak")
	assertConvert(t, c, "あかn", "akan")

	assertConvert(t, c, "かんな", "kannna")
	assertConvert(t, c, "かんあ", "kanna")
	assertConvert(t, c, "かな", "kana")

	assertConvert(t, c, "いっかん", "ikkann")
	assertConvert(t, c, "いっかn", "ikkan")
	assertConvert(t, c, "いかん", "ikann")
	assertConvert(t, c, "いかn", "ikan")

	assertConvert(t, c, "いっきん", "ikkinn")
	assertConvert(t, c, "いっきn", "ikkin")
	assertConvert(t, c, "いきん", "ikinn")
	assertConvert(t, c, "いきn", "ikin")
}

func assertUnescape(t *testing.T, text, expected string) {
	actual := unescape(text)
	if actual != expected {
		t.Errorf("unescape failed: expected=%s actual=%s", expected, actual)
	}
}

func TestUnescape(t *testing.T) {
	assertUnescape(t, "abc", "abc")
	assertUnescape(t, "a\\bc", "abc")
	assertUnescape(t, "\\\\", "\\")
	assertUnescape(t, "\\#", "#")
	assertUnescape(t, "\\", "\\")
}
