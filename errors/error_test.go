package errors

import (
	"fmt"
	"testing"
)

func TestErrorVariableDemo(t *testing.T) {
	err := WebCall(true)

	if err != nil {
		fmt.Println("Bad request happened.")
	}
}

func TestCustomErrorDemo(t *testing.T) {
	CustomErrorDemo()
}
