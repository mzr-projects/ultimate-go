package data_structures

import (
	"fmt"
	"sort"
)

type user struct {
	name    string
	surname string
}

type player struct {
	name  string
	score int
}

func MapUtils() {

	/*
		Declare and make a map that stores values of type user with a key of type string
	*/
	users := make(map[string]user)
	users["Ben"] = user{"Ben", "Ben"}
	users["James"] = user{"James", "James"}
	users["Sina"] = user{"Sina", "Sina"}
	users["Jorden"] = user{"Jorden", "Jorden"}

	sina := users["Sina"]
	fmt.Printf("%+v:\n", sina)

	/*
		Here we changed the value of Ben to Bob
	*/
	users["Ben"] = user{"Bob", "Bob"}
	fmt.Printf("%+v:\n", users["Ben"])

	/*
		Here we removed Ben from the Map
	*/
	delete(users, "Ben")
	fmt.Printf("%+v:,size: %d\n", users, len(users))

	/*
		If the key doesn't exist, we get the zero value out of it
	*/
	unKnown1 := users["Jimmy"]
	fmt.Printf("non-existence user: %+v:\n", unKnown1)

	/*
		This 2 value shows the key exists or not, ok=true means it does exist
	*/
	unKnown2, ok := users["Jimmy"]
	fmt.Printf("%+v,ok: %v\n", unKnown2, ok)

	/*
		Iteration over maps(keys and values) is like below
	*/
	for key, value := range users {
		fmt.Printf("key: %s, value: %+v\n", key, value)
	}

	/*
		Iterates over keys
	*/
	for key := range users {
		fmt.Printf("key: %s\n", key)
	}

	fmt.Println("================= Sorted keys")

	var keys []string
	for k := range users {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("key: %s\n", key)
	}
}

func PointersAndMaps() {
	players := map[string]player{
		"jimmy":  {"jimmy", 32},
		"maziar": {"maziar", 31},
	}

	/*
		Go’s maps may grow or rehash internally, causing the locations of their key-value pairs to move around.
		If the language allowed you to take a pointer into a map, that pointer could become invalid
		as soon as the map resizes or re-hashes
	*/
	//jimmy := &players["jimmy"]
	//jimmy.score++

	player := players["jimmy"]
	player.score++
	players["jimmy"] = player
	fmt.Printf("%+v\n", player)
}

func PointerSemanticAndMaps() {
	scoresTemp := map[string]int{
		"jimmy":  32,
		"maziar": 31,
	}

	/*
		Maps are reference type we pass them like a value semantic we don't pass them like *scoreTemp
		Here every change in double function will affect the original map scoresTemp, and we can see this in the
		following Printf
	*/
	double(scoresTemp, "jimmy")
	fmt.Printf("%+v\n", scoresTemp)
}

func double(temp map[string]int, s string) {
	temp[s] = temp[s] * 2
}
