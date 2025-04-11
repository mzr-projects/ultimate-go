package composition

import "fmt"

type Mover interface {
	Move()
}

type Locker interface {
	Lock()
	Unlock()
}

/*
MoveLocker is the composition for two other interfaces
*/
type MoveLocker interface {
	Mover
	Locker
}

type bike struct {
}

func (bike) Move() {
	fmt.Println("Moving bike")
}

func (bike) Lock() {
	fmt.Println("Locking bike")
}

func (bike) Unlock() {
	fmt.Println("Unlocking bike")
}

func BikeDemo() {
	/*
		Declare variables of MoveLocker and Mover interfaces set to their zero values
	*/
	var ml MoveLocker
	var m Mover

	/*
		Creat a value of type bike and assign the value to the MoveLocker interface value.
	*/
	ml = bike{}

	/*
		An interface of MoveLocker can be implicitly converted to into a value of type Mover, they both declare a method
		named Move
	*/
	m = ml

	/*
		Compile time Error, can't use m(type mover) as type MoveLocker in assignment
		Mover does not implement MoveLocker(missing Lock method)
	*/
	//ml = m

	/*
		    We have a type assertion:
			A built-in language feature that allows you to extract the concrete type and value from an interface.
	*/
	b, ok := m.(bike)
	if ok {
		fmt.Println("We have a bike")
	}
	ml = b
}
