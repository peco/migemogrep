package inflator

type suffixer struct {
	suffixes []string
}

func Suffix(suffixes ...string) Inflatable {
	return &suffixer{suffixes}
}

func (p *suffixer) Inflate(s string) <-chan string {
	return Start(func(c chan<- string) {
		for _, t := range p.suffixes {
			c <- s + t
		}
	})
}
