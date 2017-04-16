package migemo

import (
	"errors"
	"github.com/koron/gomigemo/conv"
	skkdict "github.com/koron/gomigemo/dict"
	"github.com/koron/gomigemo/inflator"
	"io"
)

type dict struct {
	assets   Assets
	inflator inflator.Inflatable
}

func (d *dict) Matcher(s string) (Matcher, error) {
	return newMatcher(d, s)
}

func (d *dict) loadSKKDict(name string) (sd *skkdict.Dict, err error) {
	err = d.assets.Get(name, func(rd io.Reader) (err error) {
		sd, err = skkdict.ReadSKK(rd)
		return err
	})
	if err != nil {
		sd = nil
	}
	return sd, err
}

func (d *dict) loadConv(name string) (c *conv.Converter, err error) {
	c = conv.New()
	err = d.assets.Get(name, func(rd io.Reader) error {
		_, err := c.Load(rd, name)
		return err
	})
	if err != nil {
		c = nil
	}
	return c, err
}

func (d *dict) load() error {
	if d.inflator != nil {
		return errors.New("Dictionaries were loaded already.")
	}

	// Load dictionaries.
	skk, err := d.loadSKKDict("SKK-JISYO.utf-8.L")
	if err != nil {
		return err
	}
	roma2hira, err := d.loadConv("roma2hira.txt")
	if err != nil {
		return err
	}
	hira2kata, err := d.loadConv("hira2kata.txt")
	if err != nil {
		return err
	}
	wide2narrow, err := d.loadConv("wide2narrow.txt")
	if err != nil {
		return err
	}

	// Build inflator.
	d.inflator = inflator.Join(
		inflator.DispatchEcho(
			inflator.Join(
				roma2hira,
				inflator.DispatchEcho(inflator.Join(
					hira2kata,
					inflator.DispatchEcho(wide2narrow),
				)),
			),
		),
		inflator.DispatchEcho(skk),
	)

	// FIXME: Make these (loader and builder) flexible.
	return nil
}
