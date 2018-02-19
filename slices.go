package main

import (
	"fmt"
)

func Slices() {
	someSlice()
}

func someSlice() {
	a := [10]int{1, 2, 3, 4, 5}
	var s []int
	s = a[2:5]

	fmt.Println("slice is ", s)
}
