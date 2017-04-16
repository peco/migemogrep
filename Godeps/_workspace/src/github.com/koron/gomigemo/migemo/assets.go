package migemo

import (
	"io"
	"os"
	"path/filepath"
)

type AssetProc func(io.Reader) error

type Assets interface {
	Get(name string, proc AssetProc) error
}

type PathAssets struct {
	root string
}

func (a *PathAssets) Get(name string, proc AssetProc) error {
	path := filepath.Join(a.root, name)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return proc(file)
}
