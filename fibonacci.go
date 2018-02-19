package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	s := []int{0, 1}
	var i int
	inc := func() {
		i++
	}
	return func() int {
		defer inc()
		if i >= 2 {
			s = append(s, s[i-2]+s[i-1])
		}
		return s[i]
	}
}

func Fibonacci() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
