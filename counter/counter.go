package counter

type Operation int

const (
	GET Operation = iota
)

type Counter struct {
	ops   chan Operation
	res   chan int
	count int
}

func New() *Counter {
	counter := Counter{
		ops: make(chan Operation, 1000000),
		res: make(chan int, 1),
	}

	go func() {
		for a := range counter.ops {
			switch a {
			case GET:
				counter.res <- counter.count
				counter.count++
			}
		}
	}()

	return &counter
}

func (c *Counter) Get() int {
	c.ops <- GET
	return <-c.res
}
