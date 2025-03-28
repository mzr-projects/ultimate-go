package composition

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

type Data struct {
	Line string
}

/*
Xenia Represents the old database that we're going to pull data from
*/
type Xenia struct {
	Host    string
	Timeout time.Duration
}

type Pillar struct {
	Host    string
	Timeout time.Duration
}

/*
The System is going to wrap Xenia and Pillar together into a single system type
*/
type System struct {
	Puller
	Storer
}

/*
Puller we discovered this interface after writing the basic code that works we didn't design the code based on this
interface we discovered it
*/
type Puller interface {
	Pull(d *Data) error
}

/*
Storer we discovered this interface after writing the basic code that works we didn't design the code based on this
interface we discovered it
*/
type Storer interface {
	Store(d *Data) error
}

type PullerStorer interface {
	Puller
	Storer
}

/*
Pull knows how to pull data from Xenia
*/
func (*Xenia) Pull(d *Data) error {

	switch rand.Intn(10) {
	case 1, 9:
		return io.EOF
	case 5:
		return errors.New("error reading data")
	default:
		d.Line = "Data"
		fmt.Println("In:", d.Line)
		return nil
	}
}

/*
Store will get the data pulled from Xenia and store in Pillar
*/
func (*Pillar) Store(d *Data) error {
	fmt.Println("Out:", d.Line)
	return nil
}

/*
pull knows how to pull a batch of data from Xenia
*/
func pull(p Puller, data []Data) (int, error) {
	for i := range data {
		if err := p.Pull(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

func store(s Storer, data []Data) (int, error) {
	for i := range data {
		if err := s.Store(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

func Copy(ps PullerStorer, batch int) error {
	data := make([]Data, batch)

	/*
		This is an Endless loop
	*/
	for {
		i, err := pull(ps, data)
		if i > 0 {
			if _, err := store(ps, data[:i]); err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}
	}
}

func DemoThree() {
	sys := System{
		Puller: &Xenia{
			Host:    "192.168.1.1",
			Timeout: time.Second,
		},
		Storer: &Pillar{
			Host:    "192.168.1.1",
			Timeout: time.Second,
		},
	}

	if err := Copy(&sys, 3); err != nil {
		fmt.Println(err)
	}
}
