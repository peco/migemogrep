package inflator

type echo struct {
}

func Echo() Inflatable {
	return &echo{}
}

func (e *echo) Inflate(s string) <-chan string {
	return Start(func(c chan<- string) {
		c <- s
	})
}
