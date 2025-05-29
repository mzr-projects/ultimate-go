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
Here the method argument says I dont want an int value, I want an address of int value
*/
func incrementPointerSemantic(inc *int) {
	*inc++
	println("inc:\tValue of[", inc, "]\tAddr of [", &inc, "]\tValue points to [", *inc, "]")
}
