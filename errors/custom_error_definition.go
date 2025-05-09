package errors

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	name string
}

type CustomError struct {
	Value string
	Type  reflect.Type
}

/*
We are implementing the Error interface, notice that we are using pointer semantics
*/
func (e *CustomError) Error() string {
	return "demo error is :" + e.Value
}

func CustomErrorDemo() {
	var user User
	err := json.Unmarshal([]byte(`{"name":"john"}`), &user)

	if err != nil {
		switch e := err.(type) {
		case *CustomError:
			fmt.Println("custom error is :" + e.Error())
		case *json.UnmarshalTypeError:
			fmt.Println("unmarshal type error :" + e.Error())
		case *json.InvalidUnmarshalError:
			fmt.Println("invalid unmarshal error :" + e.Error())
		}
	}
}
