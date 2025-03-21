package decoupling

import "fmt"

/*
Reader is an interface that defines the act of reading data.
*/
type reader interface {
	read(b []byte) (int, error)
}

/*
File defines a system file
*/
type file struct {
	name string
}

type pipe struct {
	name string
}

/*
The Definition of the method matches the exact definition of the method we defined in the reader interface
Now this concrete type will implement the read method of the reader interface under the hood
*/
func (file) read(b []byte) (int, error) {
	s := "<rss><channel><title>Go Programing Demo</title></channel></rss>"
	copy(b, s)
	return len(s), nil
}

/*
The same will happen for this method exactly like the previous one
*/
func (pipe) read(b []byte) (int, error) {
	s := `{name:"MT",age:12}`
	copy(b, s)
	return len(s), nil
}

func Retriever(r reader) error {
	data := make([]byte, 1024)

	length, err := r.read(data)
	if err != nil {
		return err
	}

	fmt.Println(string(data[:length]))
	return nil
}
