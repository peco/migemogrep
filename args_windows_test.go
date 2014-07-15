package main

import (
	"os"
	"reflect"
	"sort"
	"testing"
)

func TestExpandArgs(t *testing.T) {
	args := os.Args
	defer func() {
		os.Args = args
	}()

	os.Args = []string{"foo", "*_test.go"}
	expandArgs()

	given := os.Args
	expect := []string{"foo", "args_windows_test.go", "grep_test.go"}

	sort.Strings(given)
	sort.Strings(expect)
	if !reflect.DeepEqual(given, expect) {
		t.Fatalf("Should be %v but %v", expect, os.Args)
	}
}

func TestExpandArgsDashDash(t *testing.T) {
	args := os.Args
	defer func() {
		os.Args = args
	}()

	os.Args = []string{"foo", "--", "*_test.go"}
	expandArgs()

	expect := []string{"foo", "*_test.go"}
	if os.Args[0] != "foo" || os.Args[1] != "*_test.go" {
		t.Fatalf("Should be %v but %v", expect, os.Args)
	}
}
