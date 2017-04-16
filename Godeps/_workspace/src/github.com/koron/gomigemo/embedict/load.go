package embedict

import (
	"github.com/koron/gomigemo/migemo"
)

func Load() (migemo.Dict, error) {
	return migemo.LoadAssets(&assets{})
}
