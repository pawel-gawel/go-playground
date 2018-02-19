package main

import (
	"fmt"
)

func Arrays() {
	createArray()
	arrayOfRefs()
}

func createArray() {
	var a [10]int
	a[3] = 7
	fmt.Println("array a length is ", len(a))
	fmt.Println("array a is ", a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr with defined elements ", b)
}

func arrayOfRefs() {
	type Block struct {
		a int
	}

	var a [5]*Block
	a[3] = &Block{1}
	fmt.Println("references array is ", a)
}
