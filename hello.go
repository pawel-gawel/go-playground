package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

func main() {
	fmt.Printf("hello ziom!\n")
	fmt.Println("Random number", rand.Intn(200))
	fmt.Println("Sqrt", math.Sqrt(7))

	x, y, sum := add(3, 4)
	fmt.Printf("Add %d + %d = %f \n", x, y, sum)

	name, surname := splitName("John Doe")
	fmt.Printf("Name is %s and surname %s\n", name, surname)
}

func add(x, y int) (int, int, float32) {
	return x, y, float32(x + y)
}

func splitName(name string) (firstName, lastName string) {
	firstName = strings.Split(name, " ")[0]
	lastName = strings.Split(name, " ")[1]
	return
}
