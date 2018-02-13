package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x, y int) float32 {
	return float32(x + y)
}

func main() {
	fmt.Printf("hello ziom!\n")
	fmt.Println("Random number", rand.Intn(200))
	fmt.Println("Sqrt", math.Sqrt(7))

	fmt.Println("Add x + y = ", add(3, 4))
}
