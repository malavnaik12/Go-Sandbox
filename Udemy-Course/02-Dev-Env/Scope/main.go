/*
THIS IS ONLY SOME OF THE CODE SHOWN IN THE LECTURE
ALL OF THE CODE SHOWN IN THE LECTURE IS AVAILABLE ON
https://github.com/GoesToEleven/learn-to-code-go-version-03
look in the "031-scope" directory

THE REASON:
golang playground doesn't allow you to create multiple files / packages

                     ¯\_(ツ)_/¯
*/
// ----------------------------------------------
// --------- below is most of the code ----------
// ----------------------------------------------

package main

// the imported package "fmt"
// is in the "file block" scope
// of this file
import (
	"fmt"

	"github.com/malavnaik12/Go-Sandbox/Udemy-Course/02-Dev-Env/Scope/furtherexplored"
)

// x is in "package block" scope
var x = 42

func main() {
	// x can be accessed here
	fmt.Println(x)

	// x can be accessed here
	printMe()

	// y does not exist here yet
	// so this won't work
	// fmt.Println(y)

	// y is in "block scope"
	y := 43
	fmt.Println(y)

	p1 := person{
		"Malav",
	}
	p1.sayHello()

	// variable "shadowing"
	x := 32
	fmt.Println(x)

	// package block x is still the same
	printMe()

	//also good to know

	if z := 82; z > 50 {
		fmt.Println(z)
	}
	// you can't access z here
	// fmt.Println(z)
	/*
		take a look at the "comma ok idiom"
		https://go.dev/doc/effective_go#maps
	*/

	// Assessing inherent packages
	furtherexplored.Fascinating()
}

func printMe() {
	//x can be accessed here
	fmt.Println(x)
}

// type person is in "package block" scope
type person struct {
	first string
}

// the method sayHello
// which is attached to VALUES of TYPE person
// is in "package block" scope
func (p person) sayHello() {
	fmt.Println("Hi, my name is", p.first)
}
