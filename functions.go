package main

import "fmt"

func Functions() {
	anonymousFuncAsValue()
	closures()
}

func compute(a, b float64, fn func(float64, float64) float64) float64 {
	return fn(a, b)
}

func anonymousFuncAsValue() {
	addition := func(a, b float64) float64 {
		return a + b
	}
	res := compute(3, 4, addition)
	fmt.Println("computation is", res)
}

func closures() {
	x := 1
	fn := func(a float64) float64 {
		return float64(x) + a
	}

	fmt.Println("closured sum is ", fn(5.6))
}
