/*
Need a func main inside a package main
- The package is like a folder in which we organize things in a folder
- Package main is a way of organizing code in "main"

Ex: "fmt" is a package in the standard Go library
*/
package main

import "fmt"

func main() {
	fmt.Println("Hello World! 😊")
	const name, age = "Malav", 26
	fmt.Printf("%s is %d years old.\n", name, age)
	fmt.Printf(`This is a back string literal: %s is %d years old.
`, name, age)
	fmt.Printf("%s is %d years old.\tand the type is %T and %T.\n", name, age, name, age)
}
