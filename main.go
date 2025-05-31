package main

import (
	"ultimate-go/data-structures"
	"ultimate-go/data-types"
	"ultimate-go/semantics"
	"ultimate-go/stack"
)

func main() {
	data_types.Constants()

	semantics.PointerReview()

	semantics.DataSemantics()
	semantics.ValueSemantics()
	semantics.PointerSemantics()

	stack.StackGrow()

	data_structures.SliceDemo()
	data_structures.AppendToSlice()
}
