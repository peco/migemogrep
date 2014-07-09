package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/koron/gomigemo/migemo"
)

// Looks for possible dictionary directories.
// We should expand this to more unix-y locations, too
//
// Locations to search:
// 1. User-supplied location on the command line
//    -d "dictdir"
// 2. User-supplied location in the env var
//    GOMIGEMO_DICTDIR
//    GMIGEMO_DICTDIR (for backwards compatibility)
// 3. User-supplied config dir
//    ~/.config/gomigemo/dict
// 4. Places that sound be system config (XXX won't implement for now)
//    /var/lib/gomigemo/dict
// 5. Somewhere in the gopath where gomigemo resides
//    $GOPATH/github.com/koron/gomigemo/_dict
// 6. Current directory
func dictdir() string {
	// I want to change this to G*O*MIGEMO_DICTDIR
	d := os.Getenv("GMIGEMO_DICTDIR")
	if d != "" {
		return d
	}
	d = os.Getenv("GOPATH")
	if d == "" {
		d = "."
	}
	for _, p := range strings.Split(d, string(filepath.ListSeparator)) {
		candidate := filepath.Join(p, "src", "github.com", "koron", "gomigemo", "_dict")
		if _, err := os.Stat(candidate); err != nil {
			continue
		}
		return candidate
	}

	// fallback to current directory
	return d
}

var dictPath = flag.String("d", dictdir(), "Location to dictionary")

// Does the grepping
func grep(r io.Reader, re *regexp.Regexp) error {
	buf := bufio.NewReader(r)
	for {
		b, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		line := string(b)
		if re.MatchString(line) {
			fmt.Println(line)
		}
	}
	return nil
}

func main() {
	st := _main()
	os.Exit(st)
}

func _main() int {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage of main: [pattern] [files...]")
	}

	flag.Parse()

	if flag.NArg() != 1 && flag.NArg() != 2 {
		flag.Usage()
		flag.PrintDefaults()
		return 1
	}

	dict, err := migemo.Load(*dictPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	re, err := migemo.Compile(dict, flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}


	// If there's only one arg, then we need to match against the input
	if flag.NArg() == 1 {
		if err = grep(os.Stdin, re); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}

		// We got here, we're fine.
		return 0
	}

	// More than one arg. We must be searching against a file
	for _, arg := range flag.Args()[1:] {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		defer f.Close()
		if err = grep(f, re); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
	}

	return 0
}