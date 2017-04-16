package migemo

import (
	"errors"
	"github.com/koron/gelatin/trie"
)

type matcher struct {
	options   MatcherOptions
	trie      *trie.TernaryTrie
	pattern   string
	patterned bool
}

func newMatcher(d *dict, s string) (*matcher, error) {
	if d.inflator == nil {
		return nil, errors.New("Dictionary is not loaded")
	}
	m := &matcher{
		options: defaultMatcherOptions,
		trie:    trie.NewTernaryTrie(),
	}
	// Inflate s word, add those to trie.
	ch := d.inflator.Inflate(s)
	for w := range ch {
		m.add(w)
	}
	m.trie.Balance()
	return m, nil
}

func (m *matcher) Match(s string) (chan Match, error) {
	// FIXME: Make own match with trie in future.
	return nil, nil
}

func (m *matcher) SetOptions(o MatcherOptions) {
	m.options = o
	return
}

func (m *matcher) GetOptions() MatcherOptions {
	return m.options
}

func (m *matcher) add(s string) {
	// Add a string to m.trie.
	if len(s) == 0 {
		return
	}
	n := m.trie.Root()
	for _, c := range s {
		n, _ = n.Dig(c)
		if n.Value() != nil {
			return
		}
	}
	n.SetValue(true)
	n.RemoveAll()
}
