package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	basics()
	multipleReturns()
	variablesDeclaredInBlock()
	constants()
	forLoop()
	forIsWhile()
	foreverLoop()
	ifWithShortStatement()
	switchStatement()
	noConditionSwitch()
	deferStatement()
	Structs()
	Arrays()
}

func basics() {
	fmt.Printf("hello %s!\n", "ziom")
	fmt.Println("Random number", rand.Intn(200))
	fmt.Println("Sqrt", math.Sqrt(7))
}

func multipleReturns() {
	x, y, sum := add(3, 4)
	fmt.Printf("Add %d + %d = %f \n", x, y, sum)

	var name, surname string = "Jan", "Kowalski"
	name, surname = splitName("John Doe")
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

func variablesDeclaredInBlock() {
	var (
		str string
		bl  bool
		a   int
		b           = 1<<24 - 1
		c   float32 = 7.4
	)

	fmt.Println(str, bl, a, b, c)
}

func constants() {
	const truth bool = true
	const lie = false

	fmt.Println(truth, lie)
}

func forLoop() {
	var sum int
	for i := 0; i < 3; i++ {
		sum += i
		fmt.Printf("for loop - number is %d\n", i)
	}
	fmt.Printf("for loop - sum is %d\n", sum)
}

func forIsWhile() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Printf("for/while loop - sum is %d\n", sum)
}

func foreverLoop() {
	i := 0
	for {
		if i > 10 {
			break
		}
		i++
	}
}

func ifWithShortStatement() {
	if resolved := true; resolved {
		fmt.Println("Is resolved inside if")
	}
	// fmt.Println("Is resolved outside if? ", resolved)
}

func switchStatement() {
	switch os := "osx"; os {
	case caseCanBeAnyExpression():
		fmt.Println("switch - first case")
		fallthrough
	case "windows":
		fmt.Println("switch - falls through to lower case")
	}
}

func caseCanBeAnyExpression() string {
	return "osx"
}

func noConditionSwitch() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func deferStatement() {
	v := 0
	defer fmt.Println("world", v)
	defer fmt.Print(" whole ")
	v = 1
	fmt.Print("hello")
}
