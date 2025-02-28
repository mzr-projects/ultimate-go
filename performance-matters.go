package main

type User struct {
	name  string
	email string
}

func CreateUserV1() User {
	u := User{
		name:  "John Doe",
		email: "johndoe@gmail.com",
	}

	println("V1", &u)

	/**
	Here this function is going to return its own copy of user value, so this is a value semantic function
	*/
	return u
}

func CreateUserV2() *User {

	/*
		If we create an address of user here u := &user .... , we are going to lose our readability
		Because this function now returns the u if we stick to the following code as soon as we see the
		return &u we can tell we're sharing the u upwards to the main, and we may have problem because we are using
		pointer semantic which here compiler well share u in heap memory and Garbage collector should act now,
		so we may lose performance.
	*/
	u := User{
		name:  "John Doe",
		email: "johndoe@gmail.com",
	}

	println("V2", &u)

	return &u
}
