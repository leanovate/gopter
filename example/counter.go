package example

type Counter struct {
	n int
}

func (c *Counter) Inc() {
	c.n++
}

func (c *Counter) Dec() {
	if c.n > 3 {
		// Intentional error
		c.n -= 2
	} else {
		c.n--
	}
}

func (c *Counter) Get() int {
	return c.n
}

func (c *Counter) Reset() {
	c.n = 0
}
