package main

import (
	"fmt"
	"math"
)

func add(x, y int) int {
	return (x + y)
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	// Go Tour Step 1-3
	fmt.Printf("Hello, Everyone. My favourite number is %g\n", math.Pi)

	// Go Tour Step 4-7
	fmt.Println((add(9, 10)))
	assessAdd()

	a, b := swap("My name", "Jeff")
	fmt.Println(a, b)

	fmt.Print("A return without explicitly specifying returned variable??? ")
	fmt.Println(split(21))

	// Go Tour Step 8-10
}

func assessAdd() {
	fmt.Println("What's 9+10? 21")
}
