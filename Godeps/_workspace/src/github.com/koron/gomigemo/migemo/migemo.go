package migemo

import (
	"regexp"
)

type Dict interface {
	Matcher(string) (Matcher, error)
}

type Matcher interface {
	Match(string) (chan Match, error)
	Pattern() (string, error)
	SetOptions(MatcherOptions)
	GetOptions() MatcherOptions
}

type MatcherOptions struct {
	OpOr                  string
	OpGroupIn, OpGroupOut string
	OpClassIn, OpClassOut string
	OpWSpaces             string
	//MetaChars string
}

type Match struct {
	Start, End int
}

func Load(path string) (Dict, error) {
	return LoadAssets(&PathAssets{root: path})
}

func LoadAssets(assets Assets) (Dict, error) {
	d := &dict{assets: assets}
	err := d.load()
	if err != nil {
		return nil, err
	}
	return d, nil
}

func Compile(d Dict, s string) (*regexp.Regexp, error) {
	m, err := d.Matcher(s)
	if err != nil {
		return nil, err
	}
	return NewRegexp(m)
}

func NewRegexp(m Matcher) (*regexp.Regexp, error) {
	p, err := m.Pattern()
	if err != nil {
		return nil, err
	}
	return regexp.Compile(p)
}

func Pattern(d Dict, s string) (string, error) {
	m, err := d.Matcher(s)
	if err != nil {
		return "", err
	}
	p, err := m.Pattern()
	if err != nil {
		return "", err
	}
	return p, nil
}
