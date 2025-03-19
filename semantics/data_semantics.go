package semantics

import "fmt"

func DataSemantics() {

	fmt.Println("############## Value semantics")

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
		Here in this loop fruit variable is a copy of every element in fruits array which means it has a copy of the
		pointer every element points at and the size of every element(2 words string)
	*/
	for i, v := range fruits {
		/*
			This print function also has its own copy of the fruits elements
		*/
		fmt.Println(i, v)
	}

	/*
		    This is the pointer semantic of the range loop here we don't have any fruit variable that copy the fruits elements
				we're just accessing by index in the shared mode.

			for i := range fruits {
					fmt.Println(i, fruits[i])
				}
	*/
}

func PointerSemantics() {

	fmt.Println("############## Pointer semantics")

	friends := [5]string{"Annie", "James", "John", "Alice", "Donald"}
	fmt.Printf("Bfr[%s] -> ", friends[1])

	for i := range friends {
		friends[1] = "Jack"

		if i == 1 {
			fmt.Printf("Aft[%s]\n", friends[1])
		}
	}
}

func ValueSemantics() {

	fmt.Println("############## Value semantics")

	friends := [5]string{"Annie", "James", "John", "Alice", "Donald"}
	fmt.Printf("Bfr[%s] -> ", friends[1])

	for i, v := range friends {
		friends[1] = "Jack"

		if i == 1 {
			fmt.Printf("Aft[%s]\n", v)
		}
	}
}
