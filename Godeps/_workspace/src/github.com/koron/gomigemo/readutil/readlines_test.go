package readutil

import (
	"io"
	"testing"
)

func open(path string, t *testing.T) <-chan string {
	ch := make(chan string, 1)
	go ReadFileLines(path, func(line string, err error) error {
		if len(line) > 0 {
			ch <- line
		}
		if err != nil {
			close(ch)
			if err != io.EOF {
				t.Error(err)
			}
		}
		return err
	})
	return ch
}

func TestReadLines(t *testing.T) {
	ch := open("./readlines_test0.txt", t)
	if s1 := <-ch; s1 != "foo\n" {
		t.Error("1st line is not \"foo\\n\":", s1)
	}
	if s2 := <-ch; s2 != "bar\n" {
		t.Error("1st line is not \"bar\\n\":", s2)
	}
	if s3 := <-ch; s3 != "baz\n" {
		t.Error("1st line is not \"baz\\n\":", s3)
	}
	if s4, ok := <-ch; ok != false {
		t.Error("mode data:", s4)
	}
}

func TestReadLinesWithoutEOL(t *testing.T) {
	ch := open("./readlines_test1.txt", t)
	if s1 := <-ch; s1 != "foo\n" {
		t.Error("1st line is not \"foo\\n\":", s1)
	}
	if s2 := <-ch; s2 != "bar\n" {
		t.Error("1st line is not \"bar\\n\":", s2)
	}
	if s3 := <-ch; s3 != "baz" {
		t.Error("1st line is not \"baz\":", s3)
	}
	if s4, ok := <-ch; ok != false {
		t.Error("mode data:", s4)
	}
}
