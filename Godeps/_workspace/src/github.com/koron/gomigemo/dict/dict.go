package dict

import (
	"github.com/koron/gelatin/trie"
	"github.com/koron/gomigemo/inflator"
)

type Dict struct {
	trie     *trie.TernaryTrie
	balanced bool
}

type entry struct {
	words []string
}

func New() *Dict {
	return &Dict{
		trie:     trie.NewTernaryTrie(),
		balanced: false,
	}
}

func (d *Dict) Add(label string, words []string) {
	d.trie.Put(label, &entry{words: words})
	d.balanced = false
}

func (d *Dict) Balance() {
	if !d.balanced {
		d.trie.Balance()
		d.balanced = true
	}
}

func (d *Dict) Get(label string, proc func(word string) bool) {
	n := d.trie.Get(label)
	if n == nil {
		return
	}
	f := func(o trie.Node) bool {
		e, ok := o.Value().(*entry)
		if !ok {
			return true
		}
		for _, w := range e.words {
			if !proc(w) {
				return false
			}
		}
		return true
	}
	if !f(n) {
		return
	}
	n.Each(f)
}

func (d *Dict) GetAll(label string, max int) []string {
	limit := max
	if limit == 0 {
		limit = 32
	}
	words := make([]string, 0, limit)

	d.Get(label, func(word string) bool {
		words = append(words, word)
		if max > 0 && len(words) >= max {
			return false
		}
		return true
	})
	return words
}

func (d *Dict) Inflate(s string) <-chan string {
	return inflator.Start(func(ch chan<- string) {
		d.Get(s, func(word string) bool {
			ch <- word
			return true
		})
	})
}
