package dict

import (
	"testing"
)

func assertGetAll(t *testing.T, d *Dict, label string, expected []string) {
	actual := d.GetAll(label, 0)
	if len(actual) != len(expected) {
		t.Errorf("length not match, expected=%d actual=%d",
			len(expected), len(actual))
		t.Logf("  label=%s, expected=%v", label, expected)
		return
	}
	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("item[%d]=%v is not match with %v", i, v, expected[i])
			t.Logf("  label=%s, expected=%v", label, expected)
		}
	}
}

func TestDict(t *testing.T) {
	d := New()
	d.Add("あい", []string{"愛", "藍"})
	d.Add("あき", []string{"秋", "空き"})
	d.Add("あ", []string{"亜"})
	d.Add("いき", []string{"息", "遺棄", "粋"})
	d.Add("いし", []string{"石", "医師"})

	assertGetAll(t, d, "あい", []string{"愛", "藍"})
	assertGetAll(t, d, "あき", []string{"秋", "空き"})
	assertGetAll(t, d, "あ", []string{"亜", "愛", "藍", "秋", "空き"})
	assertGetAll(t, d, "あし", []string{})
	assertGetAll(t, d, "い", []string{"息", "遺棄", "粋", "石", "医師"})
	assertGetAll(t, d, "いき", []string{"息", "遺棄", "粋"})
	assertGetAll(t, d, "いし", []string{"石", "医師"})
	assertGetAll(t, d, "いち", []string{})
}
