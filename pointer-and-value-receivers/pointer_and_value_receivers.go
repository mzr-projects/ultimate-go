package pointer_and_value_receivers

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}

func (c *Counter) Value() int {
	return c.count
}

func WrongUpdate(c Counter) {
	c.Increment()
}

func RightUpdate(c *Counter) {
	c.Increment()
}
