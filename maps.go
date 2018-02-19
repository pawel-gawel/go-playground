package main

import (
	"fmt"
)

func Maps() {
	fmt.Println("\n----- MAPS -----\n")

	makeMap()
	mapLiterals()
	mutatingMaps()
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

func mutatingMaps() {
	m := make(map[string]float32)
	m["pi"] = 3.14
	m["e"] = 2.7
	fmt.Println(m)

	delete(m, "pi")
	fmt.Println(m)

	e, okE := m["e"]
	pi, okPi := m["pi"]
	fmt.Println(e, okE, pi, okPi)

	type Some struct {
		i int
	}

	m2 := make(map[string]Some)
	some, okSome := m2["elo"]
	fmt.Println(some, okSome)
}
