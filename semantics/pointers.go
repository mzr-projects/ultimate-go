package semantics

func PointerReview() {
	count := 10
	println("total:\tValue of[", count, "]\tAddr of [", &count, "]")

	incrementValueSemantic(count)
	println("total:\tValue of[", count, "]\tAddr of [", &count, "]")

	/*
		Here we pass the address of total
	*/
	incrementPointerSemantic(&count)
}

func incrementValueSemantic(inc int) {
	inc++
	println("inc:\tValue of[", inc, "]\tAddr of [", &inc, "]")
}

/*
Here the method argument says I don't want an int value, I want an address of int value
*/
func incrementPointerSemantic(inc *int) {
	/*
		Before we dereference the pointer (using * to get the value), we need to nil check the pointer,otherwise we get
		a panic error
	*/
	if inc == nil {
		return
	} else {
		*inc++
	}
	println("inc:\tValue of[", inc, "]\tAddr of [", &inc, "]\tValue points to [", *inc, "]")
}

/*
FailedUpdate if we pass a nil pointer to this function, we can't change the value of it, we can reassign the value
if already a value is assigned to it but here there is no value assigned to f(in pointer_test file)
*/
func FailedUpdate(g *int) {
	x2 := 10
	g = &x2
}

/*
Update to change the value pointer points to we must dereference the pointer, we can't changed the value by assigning
the address of the value(like above function)
*/
func Update(px *int) {
	*px = 20
}
