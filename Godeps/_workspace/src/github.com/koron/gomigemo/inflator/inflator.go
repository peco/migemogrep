package inflator

type Inflatable interface {
	Inflate(s string) <-chan string
}

func Start(f func(chan<- string)) <-chan string {
	c := make(chan string, 1)
	go func() {
		defer close(c)
		f(c)
	}()
	return c
}
