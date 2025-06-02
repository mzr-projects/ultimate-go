package semantics

import (
	"fmt"
	"testing"
)

func TestFailedUpdate_1(t *testing.T) {
	var f *int
	FailedUpdate(f)
	fmt.Println(f)
}

/*
In this example, you start with x in the main set to 10. When you call failedUpdate, you copy the address of x into
the parameter px. Next, you declare x2 in failedUpdate, set to 20. You then point px in failedUpdate
to the address of x2. When you return to the main, the value of x is unchanged. When you call update,
you copy the address of x into px again.However, this time you change the value of what px in update points to,
the variable x in the main. When you return to main, x has been changed.
*/
func TestUpdate_1(t *testing.T) {
	x := 10
	FailedUpdate(&x)
	fmt.Printf("failedUpdate: %d\n", x)
	Update(&x)
	fmt.Printf("update: %d\n", x)
}
