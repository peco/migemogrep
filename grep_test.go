package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

func createTempfile(content string) (*os.File, error) {
	tmp := os.TempDir()
	f, err := ioutil.TempFile(tmp, "migemogrep")
	if err != nil {
		return nil, err
	}
	_, err = f.Write([]byte(content))
	if err != nil {
		return nil, err
	}
	_, err = f.Seek(0, os.SEEK_SET)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func TestEmpty(t *testing.T) {
	f, err := createTempfile(`
`)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	opt := &grepOpt{
		optNumber:   true,
		optFilename: false,
	}

	buf := new(bytes.Buffer)
	out = buf
	defer func() {
		out = os.Stdout
	}()

	err = grep(f, regexp.MustCompile("^foo"), opt)
	if err != nil {
		t.Fatal(err)
	}

	if buf.Len() > 0 {
		t.Fatal("Should be empty")
	}
}

func TestHit(t *testing.T) {
	f, err := createTempfile(`
foobar
barbaz
`)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	opt := &grepOpt{
		optNumber:   false,
		optFilename: false,
	}

	buf := new(bytes.Buffer)
	out = buf
	defer func() {
		out = os.Stdout
	}()

	err = grep(f, regexp.MustCompile("^foo"), opt)
	if err != nil {
		t.Fatal(err)
	}

	s := buf.String()
	if s != "foobar\n" {
		t.Fatalf("Should be %v but %v", `foobar`, s)
	}
}

func TestNumber(t *testing.T) {
	f, err := createTempfile(`
barbaz
foobar
`)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	opt := &grepOpt{
		optNumber:   true,
		optFilename: false,
	}

	buf := new(bytes.Buffer)
	out = buf
	defer func() {
		out = os.Stdout
	}()

	err = grep(f, regexp.MustCompile("^foo"), opt)
	if err != nil {
		t.Fatal(err)
	}

	s := buf.String()
	expect := "3:foobar\n"
	if s != expect {
		t.Fatalf("Should be %v but %v", expect, s)
	}
}
