package strings

import (
	"fmt"
	"unicode/utf8"
)

/*
StringUtil

# String is a two-word data structure

1.Pointer to the memory
2.Size of the string

Strings in Go are immutable
Strings in Go are UTF-8
*/
func StringUtil() {
	s := "Hello man!!!"

	/*
		UTFMax is 4 --> Up to 4 bytes per encoded rune.
	*/
	var buf [utf8.UTFMax]byte

	/*
		We're using the value semantic version of the range because this is the collection of bytes, and it is a built-in,
		which is a value semantic
	*/
	for i, c := range s {

		/*
			Capture the number of bytes for this RUNE
		*/
		rl := utf8.RuneLen(c)

		/*
			string_index: Calculates the slice offset for the bytes associated
		*/
		si := i + rl

		/*
			Copy of the rune from the String to our buffer
			Every array is just a slice waiting to happen, by buf[:] array buf converts to the slice
		*/
		copy(buf[:], s[i:si])

		fmt.Printf("%2d: %q; codepoint: %#6x; encoded bytes: %#v\n", i, c, c, buf[:rl])
	}
}

func RangeMechanics() {
	/*
		Value semantic of for range
	*/
	friends := []string{"Maziar", "Ben", "James", "Sina", "Jorden"}
	for _, v := range friends {

		/*
			variable[start:end]
			In the following syntax it means: the first two elements
		*/
		friends = friends[:2]
		fmt.Printf("v[%s]\n", v)
	}

	friends = []string{"Maziar", "Ben", "James", "Sina", "Jorden"}
	/*
		Pointer semantic of the for range
	*/
	for i := range friends {
		friends = friends[:2]
		fmt.Printf("v[%s]\n", friends[i])
	}
}
