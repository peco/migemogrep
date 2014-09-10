package readutil

import (
	"container/list"
	"io"
	"strings"
)

type StackableRuneReader struct {
	readers *list.List
}

func NewStackabeRuneReader() *StackableRuneReader {
	return &StackableRuneReader{list.New()}
}

func (r *StackableRuneReader) PushFront(s string) {
	if len(s) > 0 {
		r.readers.PushFront(strings.NewReader(s))
	}
}

func (r *StackableRuneReader) ReadRune() (ch rune, size int, err error) {
	for r.readers.Len() > 0 {
		front := r.readers.Front()
		curr := front.Value.(*strings.Reader)
		ch, size, err = curr.ReadRune()
		if err != io.EOF {
			return
		}
		r.readers.Remove(front)
	}
	return 0, 0, io.EOF
}
