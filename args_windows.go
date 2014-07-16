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
			candidates := []string{}
			if matches, err := filepath.Glob(arg); err == nil {
				// remove hidden files
				for _, m := range matches {
					if m[0] != '.' {
						candidates = append(candidates, m)
					}
				}
			}
			if len(candidates) > 0 {
				args = append(args, candidates...)
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
