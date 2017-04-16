package trie

import (
	"container/list"
)

type Trie interface {
	Root() Node
	Get(string) Node
	Put(string, interface{}) Node
	Size() int
}

func NewTrie() Trie {
	return NewTernaryTrie()
}

func Get(t Trie, k string) Node {
	if t == nil {
		return nil
	}
	n := t.Root()
	for _, c := range k {
		n = n.Get(c)
		if n == nil {
			return nil
		}
	}
	return n
}

func Put(t Trie, k string, v interface{}) Node {
	if t == nil {
		return nil
	}
	n := t.Root()
	for _, c := range k {
		n, _ = n.Dig(c)
	}
	n.SetValue(v)
	return n
}

func EachDepth(t Trie, proc func(Node) bool) {
	if t == nil {
		return
	}
	r := t.Root()
	var f func(Node) bool
	f = func(n Node) bool {
		n.Each(f)
		return proc(n)
	}
	r.Each(f)
}

func EachWidth(t Trie, proc func(Node) bool) {
	if t == nil {
		return
	}
	q := list.New()
	q.PushBack(t.Root())
	for q.Len() != 0 {
		f := q.Front()
		q.Remove(f)
		t := f.Value.(Node)
		if !proc(t) {
			break
		}
		t.Each(func(n Node) bool {
			q.PushBack(n)
			return true
		})
	}
}

type Node interface {
	Get(k rune) Node
	Dig(k rune) (Node, bool)
	HasChildren() bool
	Size() int
	Each(func(Node) bool)
	RemoveAll()

	Label() rune
	Value() interface{}
	SetValue(v interface{})
}

func Children(n Node) []Node {
	children := make([]Node, n.Size())
	idx := 0
	n.Each(func(n Node) bool {
		children[idx] = n
		idx++
		return true
	})
	return children
}
