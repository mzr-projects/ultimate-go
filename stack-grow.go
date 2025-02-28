package main

const size = 1024

func StackGrow() {
	s := "Hello"
	stackCopy(&s, 0, [size]int{})
}

/*
If the size is 10 then we aren't going to have new address for string s because its not going over 2k stack,
but if we change it to 1024, then we have changes os string s address two times
*/
func stackCopy(s *string, c int, a [size]int) {
	println(c, s, *s)

	c++
	if c == 10 {
		return
	}

	stackCopy(s, c, a)
}
