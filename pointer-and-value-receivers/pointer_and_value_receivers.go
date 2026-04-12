package semantics

import (
	"fmt"
	"time"
)

type Counter struct {
	total      int
	lastUpdate time.Time
}

func (c *Counter) IncrementCounterPointerSemantic() {
	c.total++
	c.lastUpdate = time.Now()
}

func (c Counter) IncrementCounterValueSematic() {
	c.total++
}

func (c Counter) InspectCounter() {
	fmt.Printf("(Inspection) total is :%d, lastUpdate is %s\n", c.total, c.lastUpdate)
}

/*
UpdateCounterByPassingValue Be aware that the rules for passing values to functions still apply. If we pass a
value type to a function and call a pointer receiver method on the passed value,we are invoking the method on a copy.
*/
func UpdateCounterByPassingValue(c Counter) {
	c.IncrementCounterPointerSemantic()
}

func UpdateCounterByPassingPointer(c *Counter) {
	c.IncrementCounterPointerSemantic()
}
