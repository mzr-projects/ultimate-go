package decoupling

import "fmt"

type data struct {
	name string
	age  int
}

func (d data) showName() {
	fmt.Println("The name is :", d.name)
}

func (d *data) setAge(age int) {
	d.age = age
	fmt.Println("The age is :", d.age)
}

func (d data) setAgeWithValueSemantic(age int) {
	d.age = age
	fmt.Println("The age is :", d.age)
}

func Demo() {
	fmt.Println("=========== Data with pointer semantic setAge")
	d := data{"jimmy", 32}
	d.showName()
	d.setAge(30)
	fmt.Println("The data is :", d)

	fmt.Println("=========== Data with value semantic setAge")
	d1 := data{"jimmy", 32}
	d1.showName()
	d1.setAgeWithValueSemantic(30)
	fmt.Println("The data is :", d1)

	fmt.Println("=========== showName method assignment to f1")
	/*
		Here showName has no () we're using assignment here, Here f1 is a two-word data structure which the first word
		points to the function showNames of the data and the second point to the data stores in data struct.
		Here we're calling showName method indirectly through the f1 method
	*/
	f1 := d.showName
	f1()

	/*
		Because the showName method is using the value semantic, we're working with the copy of data struct, so if we
		change the name like the following code, it is not going to change the original data stored in data struct.
	*/
	d.name = "james"
	f1()

	fmt.Println("=========== setAge method assignment to f2")
	/*
		Here f2 is an assignment to a pointer semantic method so there is no copy here, and we're accessing the original
		data
	*/
	f2 := d.setAge
	f2(45)
	fmt.Println(d)

	/*
		Here we can change the name as well because we're accessing the original data in data struct through pointer
		semantic of setAge method
	*/
	fmt.Println("=========== setAge method assignment to f2 (we can change the name as well)")
	d.name = "jack"
	f2(50)
	fmt.Println(d)
}
