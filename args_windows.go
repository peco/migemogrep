package main

import (
	"os"
	"path/filepath"
)

func expandArgs() {
	args := []string{os.Args[0]}
	raw := false
	for _, arg := range os.Args[1:] {
		if arg == "--" {
			raw = true
			continue
		}
		if !raw {
			if matches, err := filepath.Glob(arg); err == nil && len(matches) > 0 {
				args = append(args, matches...)
			} else {
				args = append(args, arg)
			}
		} else {
			args = append(args, arg)
		}
	}
	os.Args = args
}

func init() {
	expandArgs()
}
