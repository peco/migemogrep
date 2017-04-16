package readutil

import (
	"bufio"
	"io"
	"os"
)

type LineProc func(line string, err error) error

// ReadLines read lines from reader, and callback proc for each line.
func ReadLines(rd io.Reader, proc LineProc) error {
	r := bufio.NewReader(rd)
	for {
		line, err := r.ReadString('\n')
		err = proc(line, err)
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return err
		}
	}
}

// ReadFileLines read a path file, and callback proc for each line.
func ReadFileLines(path string, proc LineProc) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return ReadLines(file, proc)
}
