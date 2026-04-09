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

	semantics.CreateUserV1()
	semantics.CreateUserV2()

	stack.ShowStackGrow()

	data_structures.SliceDemo()
	data_structures.AppendToSlice()
}
