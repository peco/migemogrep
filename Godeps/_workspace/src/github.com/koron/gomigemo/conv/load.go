package conv

import (
	"bytes"
	"fmt"
	"github.com/koron/gomigemo/readutil"
	"io"
	"os"
	"strings"
)

func (c *Converter) LoadFile(path string) (count int, err error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return c.Load(file, path)
}

func (c *Converter) Load(rd io.Reader, name string) (count int, err error) {
	lnum := 0
	err = readutil.ReadLines(rd, func(line string, err error) error {
		lnum++
		line = strings.TrimRight(line, " \t\r\n")
		if len(line) == 0 || line[0] == '#' {
			return err
		}
		parts := strings.SplitN(line, "\t", 3)
		if parts == nil || len(parts) < 2 {
			return fmt.Errorf("Invalid format in file %s at line %d",
				name, lnum)
		}
		key := unescape(parts[0])
		emit := unescape(parts[1])
		var remain string
		if len(parts) >= 3 {
			remain = unescape(parts[2])
		}
		c.Add(key, emit, remain)
		count++
		return err
	})
	return count, err
}

func unescape(s string) string {
	if !strings.ContainsRune(s, '\\') {
		return s
	}
	b := new(bytes.Buffer)
	b.Grow(len(s))
	escape := false
	for _, r := range s {
		if escape {
			escape = false
			b.WriteRune(r)
		} else {
			if r == '\\' {
				escape = true
			} else {
				b.WriteRune(r)
			}
		}
	}
	if escape {
		b.WriteRune('\\')
	}
	return b.String()
}

func LoadFile(path string) (*Converter, error) {
	c := New()
	_, err := c.LoadFile(path)
	if err != nil {
		return nil, err
	}
	return c, nil
}
