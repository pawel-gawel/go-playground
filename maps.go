package main

import (
	"fmt"
)

func Maps() {
	fmt.Println("\n----- MAPS -----\n")

	makeMap()
}

func makeMap() {
	var m map[string]int
	m = make(map[string]int)
	m["one"] = 1

	fmt.Println(m)
}
