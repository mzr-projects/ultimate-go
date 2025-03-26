package composition

import "fmt"

/*
Animal
In Go we're not seeking what the type is we seek what the type DOES here in this code we're kinda using subtyping which
there is no meaning for subtyping in Go like object oriented languages(JAVA,and so on)
*/
type Animal struct {
	Name     string
	IsMammal bool
}

/*
Speak
It is not going to be used for Animal because there is no subtyping in GO what is going to be used is Speak in Cat and
Dog types
*/
func (a *Animal) Speak() {
	fmt.Printf("Animal %s Speaks! and its %t I am mammmal\n", a.Name, a.IsMammal)
}

type Dog struct {
	Animal
	PackFactor int
}

func (d *Dog) Speak() {
	fmt.Printf("Woof! My name is %s its %t I am mammal with pack %d\n", d.Name, d.IsMammal, d.PackFactor)
}

type Cat struct {
	Animal
	ClimbFactor int
}

func (c *Cat) Speak() {
	fmt.Printf("Meow! My name is :%s its %t I am mammmal with climb %d\n", c.Name, c.IsMammal, c.ClimbFactor)
}

func CompositionDemo() {
	/*
		Here Compiler will give us Error

		animals := []Animal{
			Dog{Animal: Animal{"Dog", true}, PackFactor: 5},
			Cat{Animal: Animal{"Cat", true}, ClimbFactor: 10},
		}*/
}
