package trie

import (
	"testing"
)

func checkTrieNode(t *testing.T, n Node, ch rune, value int) {
	if n == nil {
		t.Fatal("TrieNode is null")
	}
	if l := n.Label(); l != ch {
		t.Errorf("TrieNode.Label() expected:'%c' actual:'%c'", ch, l)
	}
	if v := n.Value().(int); v != value {
		t.Errorf("TrieNode.Value() expected:%d actual:%d", value, v)
	}
}

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Put("1", 111)
	trie.Put("2", 222)
	trie.Put("3", 333)
	trie.Put("4", 444)
	trie.Put("5", 555)

	nodes := Children(trie.Root())
	checkTrieNode(t, nodes[0], '1', 111)
	checkTrieNode(t, nodes[1], '2', 222)
	checkTrieNode(t, nodes[2], '3', 333)
	checkTrieNode(t, nodes[3], '4', 444)
	checkTrieNode(t, nodes[4], '5', 555)

	if s := trie.Size(); s != 5 {
		t.Errorf("trie.Size() returns not 5: %d", s)
	}
}

func TestNotFound(t *testing.T) {
	trie := NewTrie()
	if trie.Get("not_exist") != nil {
		t.Errorf("found 'not_exist' in empty trie")
	}
}
