package main

import "fmt"

var outside int = 42 // This definition scheme needs to be followed at the package level (outside of a function)

func main() {
	// Intro
	var h int = 42
	var hh string = "lol"
	h_new := 42
	hh_new := "lol"
	fmt.Printf("Hello, value of h is %d and hh is %s\n", h, hh)
	fmt.Printf("Similarly, Hello, value of h is %d and hh is %s\n", h_new, hh_new)

	// Printing Binary and hexadecimals
	fmt.Printf("Next, in h in binary is %b and in hexadecimals is %#X\n", h, h_new)

	// Numeral Systems: Decimal, Binary, Hexadecimals
	// No code

	//Values, types, and conversions
	y := 42 // This type of definition is only possible inside of a function
	z := 42.0

	fmt.Printf("y: %v of type %T \n", y, y)
	fmt.Printf("z: %v of type %T \n", z, z)

	var m float32 = 42.742
	fmt.Printf("m: %v of type %T \n", m, m)

	// z = m // This will not work
	z = float64(m) // This will work
	fmt.Printf("z: %v of type %T \n", z, z)

	// Built-in Types, Aggregate Types, Composition
	// Aggregate: Array, slice, struct

}
