package embedict

import (
	"bytes"
	"github.com/koron/gomigemo/migemo"
)

type assets struct {
}

func (*assets) Get(name string, proc migemo.AssetProc) error {
	b, err := Asset(name)
	if err != nil {
		return err
	}
	return proc(bytes.NewReader(b))
}
