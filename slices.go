package main

import (
	"fmt"
)

var a = [10]int{1, 2, 3, 4, 5}

func Slices() {
	someSlice()
	sliceReferencing()
	sliceLiterals()
	sliceDefaults()
	sliceCapacity()
	makeSlice()
	ticTacToe()
	slaceRange()
}

func someSlice() {
	var s []int
	s = a[2:5]

	fmt.Println("slice is ", s)
}

func sliceReferencing() {
	s1 := a[2:5]
	s2 := a[4:10]
	s1[2] = 10
	fmt.Println("modified array is ", a)
	fmt.Println("changes are seen in slice ", s2)
}

func sliceLiterals() {
	sl := []string{"a", "b", "c", "d"}
	fmt.Println("slice literal is ", sl)
	sl = append(sl, "e")

	sl2 := []struct{ a, b int }{
		{1, 2},
		{3, 4},
	}
	fmt.Println("weird slice literal syntax ", sl2)
}

func sliceDefaults() {
	s := a[:]
	fmt.Println("slice corresponding to whole array s ", s)
}

func sliceCapacity() {
	s := a[:]
	fmt.Println("slice length is ", len(s))
	fmt.Println("slice capacity is ", cap(s))

	s = append(s, 11)
	fmt.Println("slice length is ", len(s))
	fmt.Println("slice capacity is ", cap(s))
}

func makeSlice() {
	length, capacity := 5, 10
	s1 := make([]int, length)
	s2 := make([]int, length, capacity)
	fmt.Println("make slices ", s1, cap(s1), s2, cap(s2))
}

func ticTacToe() {
	s := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}
	s[1][1] = "X"
	fmt.Println(s)

}

func slaceRange() {
	var s = []int{1, 2, 3, 4}
	for i, v := range s {
		fmt.Println(i, v)
	}
}
