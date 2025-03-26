package composition

import "fmt"

/*
	Some guidelines around creating TYPEs
	1. Declare type that represents something new and unique
	2. Validate that a value of any type is created or used on its own
	3. Embed types to reuse existing behaviors you need to satisfy
	4. Question types that are an alias or abstraction for an existing type
	5. Question types whose sole purpose is to share common state
*/

/*
Speaker
Here instead of thinking about what the animal is! We think about what the animal DOES unlike the demo_one code we
designed like an OOP language
*/
type Speaker interface {
	Speak()
}

type dog struct {
	Name       string
	IsMammal   bool
	PackFactor int
}

func (d *dog) Speak() {
	fmt.Printf("Woof! My name is %s its %t I am mammal with pack %d\n", d.Name, d.IsMammal, d.PackFactor)
}

type cat struct {
	Name        string
	IsMammal    bool
	ClimbFactor int
}

func (c *cat) Speak() {
	fmt.Printf("Meow! My name is :%s its %t I am mammmal with climb %d\n", c.Name, c.IsMammal, c.ClimbFactor)
}

func SecondCompositionDemo() {
	/*
		Here we create a list of animals that know how to speak
	*/
	speakers := []Speaker{
		&dog{
			Name:       "TIM",
			IsMammal:   true,
			PackFactor: 5,
		},
		&cat{
			Name:        "Jack",
			IsMammal:    false,
			ClimbFactor: 2,
		},
	}

	for _, speaker := range speakers {
		speaker.Speak()
	}
}
