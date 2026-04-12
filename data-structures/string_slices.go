package data_structures

import (
	"fmt"
	"unicode/utf8"
)

func StringSlice() {
	// Define a slice of strings
	s := "Hello, World!"

	// UTFMax is the maximum number of bytes required to encode a rune in UTF-8, which is 4 bytes.
	var buf [utf8.UTFMax]byte

	//We are using value sematic of for range loop.
	for i, r := range s {
		//Capture the number of bytes for the current rune
		rl := utf8.RuneLen(r)

		//Calculate the slice offset for the bytes asscociated with the rune
		si := i + rl

		//Copy of rune from the string to our buffer
		copy(buf[:], s[i:si])

		fmt.Printf("%2d: %q; codepoint: %#6x; encoded bytes: %#v\n", i, r, r, buf[:rl])
	}

}
