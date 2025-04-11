package composition

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

func (u *user) String() string {
	return fmt.Sprintf("name: %s, email: %s", u.name, u.email)
}

func DemoSix() {
	user := user{
		name:  "John Doe",
		email: "johndoe@example.com",
	}

	fmt.Println(user)
	/*
		&user will run our implementation of String() cause we used the pointer semantic in the user Sting() method
	*/
	fmt.Println(&user)
}
