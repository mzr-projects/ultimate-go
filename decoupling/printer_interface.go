package decoupling

import "fmt"

type printer interface {
	print()
}

type userV2 struct {
	name string
}

func (u userV2) print() {
	fmt.Printf("Hello, %s!\n", u.name)
}

/*
Here to create our own println function, we're using empty interface which we can pass anything to it
*/
func myPrintln(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Printf("string %s\n", v)
	case int:
		fmt.Printf("int %d\n", v)
	case float64:
		fmt.Printf("float %f\n", v)
	default:
		fmt.Printf("unknown type %T\n", v)
	}
}

func PrinterDemo() {
	user := userV2{"Jordan"}

	/*
		Here we add the values and pointers to the slice of printer interface
	*/
	entities := []printer{
		//Store a copy of user value
		user,
		//Store a copy of the address of the user value interface
		&user,
	}

	user.name = "James"

	/*
		The pointer semantic of the entities slice entities[1] will see the change in the name
	*/
	for _, entity := range entities {
		entity.print()
	}

	//fmt.Printf("value: %+v,address: %v\n", entities[0], entities[1])
}

func MyPrintlnDemo() {
	s := "hello world"
	myPrintln(s)

	i := 32
	myPrintln(i)

	f := 3.14
	myPrintln(f)
}
