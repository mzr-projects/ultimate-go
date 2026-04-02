package data_structures

import (
	"fmt"
	"testing"
)

func TestSliceDemo(t *testing.T) {
	AppendToSlice()
	fmt.Println("======= Slice Demo functional test finished. =======")
	SliceDemo()
}

func TestSliceOfSlice(t *testing.T) {
	SliceOfSlice()
}

func TestSliceModification(t *testing.T) {
	SliceModification()
}
