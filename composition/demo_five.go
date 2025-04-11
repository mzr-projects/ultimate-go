package composition

import (
	"fmt"
	"math/rand"
	"time"
)

type car struct {
}

func (car) String() string {
	return "We have a car"
}

type cloud struct {
}

func (cloud) String() string {
	return "AWS"
}

func DemoFive() {
	rand.Seed(time.Now().UnixNano())

	mvs := []fmt.Stringer{
		car{},
		cloud{},
	}

	for i := 0; i < 10; i++ {
		rn := rand.Intn(2)
		if v, is := mvs[rn].(cloud); is {
			fmt.Println("Got lucky:", v)
			continue
		}
		fmt.Println("We got unlucky")
	}
}
