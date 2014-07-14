package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
)

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
				fmt.Printf("%s:", opt.filename)
			}
			if opt.optNumber {
				fmt.Printf("%d:", n)
			}
			fmt.Println(line)
		}
		n++
	}
	return nil
}

