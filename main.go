package main

import (
	"fmt"
	"time"
)

/*
Defining a new type includes some fields
this struct gets some memory: one byte for bool, two bytes for int16, four bytes for float32 so this struct gets
seven bytes of memory, but in the actual world it doesn't get seven bytes because of OS memory alignment and there
is one byte for the alignment. Because a 64bit machine carries 64 eight bytes words here we have struct with seven bytes,
so to fit this in one operation we need one byte padding to fit this in an eight-byte word

	Type example struct {
		flag bool
		counter int64
		pi float32
	}

To get rid of this padding stuff, the best way is to order our fields from largest to smallest
*/
type example struct {
	counter int64
	pi      float32
	flag    bool
}

type bill struct {
	counter int64
	pi      float32
	flag    bool
}

type nancy struct {
	counter int64
	pi      float32
	flag    bool
}

func main() {
	variables()
	time.Sleep(1 * time.Second)
	Constants()
}

func variables() {

	/*
		 Go has int16,int23,int64. If we're in 32-bit cpu then int32 in the most efficient int
		but if we're an int64 bit cpu then int64 is the most efficient int

		Here by using var, we create a zero-value variable
	*/
	var a int
	var b string
	var c float64 // Represent 64 bytes of memory
	var d bool    // 1 byte of memory

	fmt.Printf("var a int is: \t %T [%v]\n", a, a)
	fmt.Printf("var b string is: \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 is: \t %T [%v]\n", c, c)
	fmt.Printf("var d bool is: \t %T [%v]\n", d, d)

	fmt.Println("############## Initialization")

	/*
		Here we declare and initialize variables at the same time
	*/
	aa := 10
	bb := "hello"
	cc := 3.14567
	dd := true

	fmt.Printf("var aa int is: \t %T [%v]\n", aa, aa)
	fmt.Printf("var bb string is: \t %T [%v]\n", bb, bb)
	fmt.Printf("var cc float64 is: \t %T [%v]\n", cc, cc)
	fmt.Printf("var dd bool is: \t %T [%v]\n", dd, dd)

	/*
			Specify a type and perform the conversion
		 10 is an int with one byte on memory allocation here we say to convert it into a 32-bit integer which
		 allocates 4byte of memory instead of one byte
	*/
	aaa := int32(10)
	fmt.Printf("var dd bool is: \t %T [%v]\n", aaa, aaa)

	fmt.Println("############## Working with STRUCTS")

	var e1 example
	fmt.Printf("struct example is: \t %T [%+v]\n", e1, e1)

	e2 := example{
		counter: 10,
		pi:      3.14567,
		flag:    true,
	}
	fmt.Println("Flag is: ", e2.flag)
	fmt.Println("Pi is: ", e2.pi)
	fmt.Println("Counter is: ", e2.counter)

	fmt.Println("##### Anonymous Type")

	var anonymous struct {
		flag    bool
		counter int32
		pi      float32
	}
	fmt.Printf("struct anonymous: %+v\n", anonymous)

	/*
		Here we created an anonymous type and initialized it at the same time
	*/
	anonymous2 := struct {
		counter int64
		pi      float32
		flag    bool
	}{
		flag:    true,
		counter: 10,
		pi:      3.14,
	}
	fmt.Printf("struct anonymous2: %+v\n", anonymous2)

	/*
		We can't do the following, we can't assign a type bill to type nancy even they have exact same fields and memory
		allocation and padding, go doesn't support implicit conversion.

			var bi bill
			var na nancy
			bi = na
	*/

	/*
		But we can have bill(na) for explicit type conversion
	*/
	var bi bill
	var na nancy
	bi = bill(na)
	fmt.Println(bi, na)

	/*
		Here we can have conversion without being explicit because anonymous2 is not based on named type but a literal
		type
	*/
	bi = anonymous2
	fmt.Println(bi, na)

	fmt.Println("############## Pointers")
	PointerReview()

	CreateUserV1()
	CreateUserV2()

	StackGrow()
}

func PointerReview() {
	count := 10
	println("count:\tValue of[", count, "]\tAddr of [", &count, "]")

	incrementValueSemantic(count)
	println("count:\tValue of[", count, "]\tAddr of [", &count, "]")

	/*
		Here we pass the address of count
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
