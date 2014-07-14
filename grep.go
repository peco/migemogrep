package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

var out io.Writer

func init() {
	out = os.Stdout
}

// Does the grepping
func grep(r io.Reader, re *regexp.Regexp, opt *grepOpt) error {
	buf := bufio.NewReader(r)
	n := 1
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
			if opt.optFilename {
				fmt.Fprintf(out, "%s:", opt.filename)
			}
			if opt.optNumber {
				fmt.Fprintf(out, "%d:", n)
			}
			fmt.Fprintln(out, line)
		}
		n++
	}
	return nil
}

