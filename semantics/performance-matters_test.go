package semantics

import (
	"testing"
)

func TestCreateUserV1(t *testing.T) {
	CreateUserV1()
}

func TestCreateUserV2(t *testing.T) {
	CreateUserV2()
}

func TestCombined(t *testing.T) {
	u1 := CreateUserV1()
	u2 := CreateUserV2()
	println("u1: ", &u1, " u2: ", u2, "")
}
