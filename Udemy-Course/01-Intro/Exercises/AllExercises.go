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

func assessAdd() {
	fmt.Println("What's 9+10? 21")
}

// const (
// 	c0 = iota // c0 == 0
// 	c1 = iota // c1 == 1
// 	c2 = iota // c1 == 2
// )

const (
	_ = iota // c0 == 0
	aaa
	bbb
	ccc
	ddd
	eee
	fff
)

type ByteSize int

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota) // Bit shift by 10
	MB
	GB
	TB
	PB
	EB
)

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
	var c, python, java bool
	var i int
	const e, d int = 42, 11 // Error if 11 is 11.1 instead since 11.1 is float
	fmt.Println(i, c, python, java, e, d)
	fmt.Printf("%T \t %T \t %T \t %T \t %T \t \n", i, c, python, java, e)

	short_declaration_operator := "hi" // This operator is not available outside of a function
	fmt.Println(short_declaration_operator)

	// Go Tour Step 11-13
	// Various basic types like bool, string, int, byte, rune, float, complex
	fmt.Printf("Type: %T Value: %v\n", math.MaxInt, math.MaxInt)
	fmt.Printf("Converted Type: %T Value: %v\n", float64(math.MaxInt), float64(math.MaxInt))

	// Go Tour Step 14-15
	var y int = 77 // Declaring the specific type of int
	yy := 77       // Not declaring specific type, here the type is assumed depending on required level of precision

	fmt.Println(y, yy)
	fmt.Printf("y is of type %T\n", y)

	// Go Tour Step 16-17 - bitwise operations/bit shifting
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<1, 1<<1)
	fmt.Printf("%d \t %b\n", 1<<2, 1<<2)
	fmt.Printf("%d \t %b\n", 1<<3, 1<<3)
	fmt.Printf("%d \t %b\n", 1<<4, 1<<4)
	fmt.Printf("%d \t %b\n", 1<<5, 1<<5)
	fmt.Printf("%d \t %b\n", 1<<6, 1<<6)

	// Iota Integration
	fmt.Println("Iota Integration")
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<aaa, 1<<aaa)
	fmt.Printf("%d \t %b\n", 1<<bbb, 1<<bbb)
	fmt.Printf("%d \t %b\n", 1<<ccc, 1<<ccc)
	fmt.Printf("%d \t %b\n", 1<<ddd, 1<<ddd)
	fmt.Printf("%d \t %b\n", 1<<eee, 1<<eee)
	fmt.Printf("%d \t %b\n", 1<<fff, 1<<fff)

	// Measuring bits using Iota - 8 bits = 1 byte
	fmt.Println("Measuring bits using Iota")
	fmt.Printf("KB: %d \t %b\n", KB, KB)
	fmt.Printf("GB: %d \t %b\n", MB, MB)
	fmt.Printf("TB: %d \t %b\n", GB, GB)
	fmt.Printf("PB: %d \t %b\n", TB, TB)
	fmt.Printf("EB: %d \t %b\n", PB, PB)

	// Various types of items in a program
	fmt.Println("Various types of items/Variable Definitions")
	var yyy int                            // Var for zero value
	abc := "lol"                           // Short declaration operator
	var abc1, abc2 string = "lol1", "lol2" // Multiple initializations
	var cat float32 = 42.42                // Var when specificity is required
	cat1, _ := 22, 12                      // blank identifier
	fmt.Printf("Var for zero value: %d \n", yyy)
	fmt.Printf("Short declaration operator: %s \n", abc)
	fmt.Printf("Multiple initializations: %s, %s \n", abc1, abc2)
	fmt.Printf("Var when specificity is required: %v is %T\n", cat, cat)
	fmt.Printf("blank identifier: %d \n", cat1)

	// printf verbs to show values and types
	fmt.Println("printf verbs to show values and types")
	str, inter, floater := "happy", 42, 42.42
	fmt.Printf("%v is of type %T\n", str, str)
	fmt.Printf("%v is of type %T\n", inter, inter)
	fmt.Printf("%v is of type %T\n", floater, floater)

	// printf binary, decimal, hexadecimal
	fmt.Println("printf binary, decimal, hexadecimal")
	value := 12
	fmt.Printf("%v in binary: %b, decimal: %d, and hexadecimal: %#X\n", value, value, value, value)

	// signed and unsigned int
	fmt.Println("signed and unsigned int")
	var value_1 int8 = 2
	var value_2 uint8 = 129
	fmt.Printf("%v as int: %T\n", value_1, value_1)
	fmt.Printf("%v as int: %T\n", value_2, value_2)

}
