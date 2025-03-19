package data_structures

import "fmt"

/*
SliceDemo

Go has 3 different types:
 1. Built-In types
 2. Struct types
 3. Reference types (slices, maps, channels, interfaces)

The reference types and built-in type use the value semantics to move the data around
Reading and writing in reference type is going to be done in Pointer semantic
*/
func SliceDemo() {

	fmt.Println("################ Slice Demo")
	/*
		Here we create a slice with length of 5 elements

		Slices are always a 3-word data structures
		1st word Is a pointer to backing array which holds the actual 5 elements
		2nd word represents the length of the slice which is 5 here
		3rd word is the Capacity and here its set to 5
	*/
	fruits := make([]string, 5)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Pear"
	/*
		If we do this we going to get Index out of bound error
	*/
	//fruits[5] = "Apple"

	/*
		Here we're using the Value semantic to pass fruits into Printf function, it means Printf function is going
		to have its own copy of the 3-word slice(fruits)
	*/
	fmt.Printf("%v\n\n", fruits)

	/*
		Here we create slice with length 5 and capacity 8 the capacity shows the efficiency for the future growth of
		the slice, and we only can read and write those 5 elements.
	*/
	fruitsCap := make([]string, 5, 8)
	fruitsCap[0] = "Apple"
	fruitsCap[1] = "Orange"
	fruitsCap[2] = "Banana"
	fruitsCap[3] = "Grape"
	fruitsCap[4] = "Pear"
	//fruitsCap[5] = "AppleInCap6"
	//fruitsCap[6] = "OrangeInCap7"
	fmt.Printf("%v\n\n", fruitsCap)

	InspectSlice(fruits)
	InspectSlice(fruitsCap)

	AppendToSlice()
}

/*
InspectSlice

We use the value semantics to pass the slice and move slices around we shouldn't see *[]string
*/
func InspectSlice(slice []string) {
	fmt.Printf("Length[%d]  Capacity[%d]\n", len(slice), cap(slice))

	for i, s := range slice {
		fmt.Printf("[%d]\t %p %s\n", i, &slice[i], s)
	}
}

func AppendToSlice() {
	/*
		Define a nil value slice, this is still the 3-word value but the pointer is nil, the length and the capacity are
		0 as well
	*/
	var data []string

	lastCap := cap(data)
	println("lastCap:", lastCap)

	/*

	 */
	for i := 0; i < 1e5; i++ {
		value := fmt.Sprintf("Rec: %d", i)
		/*
			append will get its own copy of the data value
		*/
		data = append(data, value)
	}
}

/*
SliceOfSlice

Creating a sub-slice (for example, sub: = slice[2:5]) doesn’t allocate a new array—it simply creates a new header pointing to
a portion of the original array. This can lead to situations where the entire underlying array remains in memory because
it is still referenced by a sub-slice, even if most of its data isn’t needed.
*/
func SliceOfSlice() {
	fruitsCap := make([]string, 5, 8)
	fruitsCap[0] = "Apple"
	fruitsCap[1] = "Orange"
	fruitsCap[2] = "Banana"
	fruitsCap[3] = "Grape"
	fruitsCap[4] = "Pear"
	fmt.Println("============ Original Slice")
	InspectSlice(fruitsCap)

	/*
		Here wa are slicing the slice -> [starting_index: starting_index+length] for example [2:2+2] here length would be
		2 from index 2
	*/
	fruitsCap2 := fruitsCap[2:4]
	fmt.Println("============ Slice of Slice from 2 up to 4 not including 4")
	InspectSlice(fruitsCap2)

	//fmt.Println("============ Change the Slice")
	/*
		    The following change will mutate the fruitCap in index 0 and fruitCap2 in index 2 because they're sharing the
			same back array in the memory.
	*/
	//fruitsCap2[0] = "Gum"
	//InspectSlice(fruitsCap)
	//InspectSlice(fruitsCap2)

	fmt.Println("============ Change the Slice with Capacity Change")
	/*
		To solve the previous side effect, we cant use the following syntax of slice creation
		Here, we start from index 2 up to 4 not including 4 and the capacity from 2 up to 4 not including 4
	*/
	fruitsCap3 := fruitsCap[2:4:4]
	fruitsCap3 = append(fruitsCap3, "Gum")
	InspectSlice(fruitsCap3)
}
