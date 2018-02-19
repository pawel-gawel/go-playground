package main

import (
	"fmt"
)

func Maps() {
	fmt.Println("\n----- MAPS -----\n")

	makeMap()
	mapLiterals()
}

func makeMap() {
	var m map[string]int
	m = make(map[string]int)
	m["one"] = 1

	fmt.Println(m)
}

func mapLiterals() {
	m := map[string]int{
		"one": 1,
		"two": 2,
	}

	fmt.Println(m)
}
