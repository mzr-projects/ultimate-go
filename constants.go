package main

import (
	"fmt"
	"reflect"
)

/*
Constants only exist in compile time
*/
func Constants() {

	fmt.Println("############## Constants")

	/*
		Untyped constants
	*/
	const ui = 1234      // kind: integer
	const uf = 1234.1234 // kind: float

	fmt.Printf("%T %v, Type: %v\n", ui, ui, reflect.TypeOf(ui))
	fmt.Printf("%T %v, Type: %v\n", uf, uf, reflect.TypeOf(uf))

	/*
		Typed base constants
	*/
	const ti int = 1234          // type: int
	const tf float64 = 1234.1234 // type: float64

	//Compile time error: constant 1000 overflows uint8
	//const myUnit8 uint8 = 1000

	var answer = 3 * 0.333 //KindFloat(3) * KindFloat(0.333)
	fmt.Printf("KindFloat(3) * KindFloat(0.333): %T %v\n", answer, answer)

	const third = 1 / 3.0
	fmt.Printf("KindFloat(1) * KindFloat(3.0): %T %v\n", third, third)

	const zero = 1 / 3
	fmt.Printf("KindInt(1) * KindInt(3.0): %T %v\n", zero, zero)

	const one int8 = 1
	const two = 2 * one
	fmt.Printf("2 * one (int8(2) * int8(1)): %T %v\n", two, two)

	const (
		A1 = iota
		B1 = iota
		C1 = iota
	)
	fmt.Println("1:", A1, B1, C1)

	/*
		Here the compiler will fill the remaining iotas automatically
	*/
	const (
		A2 = iota
		B2
		C2
	)
	fmt.Println("2:", A2, B2, C2)

	const (
		A3 = iota + 1
		B3
		C3
	)
	fmt.Println("3:", A3, B3, C3)

	/*
		Here we're shifting by one at the first line and the others shift by 2,4,8,16,32
	*/
	const (
		Ldate = 1 << iota
		Ltime
		Lmicroseconds
		Llongfile
		Lshortfile
		LUTC
	)
	fmt.Println("Log:", Ldate, Ltime, Lmicroseconds, Llongfile, Lshortfile, LUTC)
}
