package main

import "fmt"

func DataSemantics() {

	/*
		Strings are two words structure each word is 4 bytes, Here we have an array of 5 string which means machine
		allocates the 40 bytes for this array.

			1st word: A pointer to a contiguous byte array containing the actual characters (UTF-8 encoded).
			2nd word: The size of the string (number of bytes it takes)

			When we have fruits[0]="Apple" we actually are copying the 2-word string "Apple" into the first index of the
			fruit string.
	*/
	var fruits [5]string
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Pear"

	/*
		This is the value semantic representation of the for range loop
	*/
	for i, v := range fruits {
		fmt.Println(i, v)
	}
}
