package main

import (
	"ultimate-go/data-structures"
	"ultimate-go/semantics"
	"ultimate-go/stack"
	"ultimate-go/variables"
)

func main() {
	variables.Constants()

	semantics.PointerReview()

	semantics.DataSemantics()
	semantics.ValueSemantics()
	semantics.PointerSemantics()

	stack.StackGrow()

	data_structures.SliceDemo()
	data_structures.AppendToSlice()
}
