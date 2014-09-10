package dict

import (
	"github.com/koron/gomigemo/skk"
	"io"
)

func addDictEntry(d *Dict, entry *skk.DictEntry) {
	words := make([]string, len(entry.Words))
	for i, w := range entry.Words {
		words[i] = w.Text
	}
	d.Add(entry.Label, words)
}

func (d *Dict) LoadSKK(path string) (count int, err error) {
	err = skk.LoadDict(path, func(entry *skk.DictEntry) {
		addDictEntry(d, entry)
		count++
	})
	return count, err
}

func LoadSKK(path string) (*Dict, error) {
	d := New()
	_, err := d.LoadSKK(path)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func ReadSKK(rd io.Reader) (d *Dict, err error) {
	d = New()
	err = skk.ReadDict(rd, func(entry *skk.DictEntry) {
		addDictEntry(d, entry)
	})
	if err != nil {
		d = nil
	}
	return d, err
}
