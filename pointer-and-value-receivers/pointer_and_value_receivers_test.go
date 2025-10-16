package pointer_and_value_receivers

import (
	"fmt"
	"testing"
)

func TestCounter_Increment(t *testing.T) {
	t.Parallel()
	var c Counter
	c.Increment()
	fmt.Printf("counter value is: %d\n", c.Value())
}

func TestWrongUpdate(t *testing.T) {
	t.Parallel()
	var c Counter
	/*
		Here we are passing the copy of c into the function so the changes will not be seen outside the scope of
		WrongUpdate, and we will get the 0 value for c
	*/
	WrongUpdate(c)
	fmt.Printf("(WrongUpdate) counter value is: %d\n", c.Value())
}

func TestRightUpdate(t *testing.T) {
	t.Parallel()
	var c Counter
	RightUpdate(&c)
	fmt.Printf("(RightUpdate) counter value is: %d\n", c.Value())
}
