package migemo

import (
	"os"
	"path/filepath"
)

var defaultMatcherOptions = MatcherOptions{
	OpOr:       "|",
	OpGroupIn:  "(?:",
	OpGroupOut: ")",
	OpClassIn:  "[",
	OpClassOut: "]",
	OpWSpaces:  "\\s*",
	// FIXME: Support MetaChars customization in future.
	//MetaChars:  "",
}

// DefaultDictdir returns default dictonary directory.
func DefaultDictdir() string {
	dir := os.Getenv("GOMIGEMO_DICTDIR")
	if dir != "" {
		return dir
	}
	// GMIGEMO_DICTDIR is obsolete
	dir0 := os.Getenv("GMIGEMO_DICTDIR")
	if dir0 != "" {
		return dir0
	}
	parts := []string{"", "src", "github.com", "koron", "gomigemo", "_dict"}
	for _, p := range filepath.SplitList(os.Getenv("GOPATH")) {
		parts[0] = p
		d := filepath.Join(parts...)
		if f, err := os.Stat(d); err == nil && f.IsDir() {
			return d
		}
	}
	// Fallback to current directory.
	return "."
}

// LoadDefault loads a dictionary with default dictdir.
func LoadDefault() (Dict, error) {
	return Load(DefaultDictdir())
}
