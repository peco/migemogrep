package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"

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
func checkdir(dir string) error {
	fi, err := os.Stat(dir)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return fmt.Errorf("not a directory")
	}
	return nil
}

// This function looks for the default dictionary location
func dictdir() string {
	dirs := []string{
		os.Getenv("GOMIGEMO_DICTDIR"),
		os.Getenv("GMIGEMO_DICDIR"),
		filepath.Join(os.Getenv("HOME"), ".config", "gomigemo", "dict"),
		filepath.Join(os.Getenv("USERPROFILE"), ".config", "gomigemo", "dict"),
		filepath.Join(os.Getenv("APPDATA"), "gomigemo", "dict"),
	}

	wd, err := os.Getwd()
	if err != nil {
		wd = ""
	}

	// Add GOPATH
	if d := os.Getenv("GOPATH"); d != "" {
		for _, p := range filepath.SplitList(d) {
			p = filepath.Join(p, "src", "github.com", "koron", "gomigemo", "_dict")
			dirs = append(dirs, p)
		}
	} else if wd != "" {
		dirs = append(dirs, filepath.Join(wd, "src", "github.com", "koron", "gomigemo", "_dict"))
	}

	for _, dir := range dirs {
		if dir == "" {
			continue
		}
		if err := checkdir(dir); err == nil {
			return dir
		}
	}

	// fallback to current directory
	return wd
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
		fmt.Fprintln(os.Stderr, "Usage: migemogrep [options] pattern [files...]")
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
