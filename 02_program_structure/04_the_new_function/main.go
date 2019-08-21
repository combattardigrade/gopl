package main

import "fmt"

func main() {
	// Another way to create a variable is to use the built-in function new.
	// The expression new(T) creates an unnamed variable of type T, and initializes
	// it to the zero value of T, and returns its address, which is a value of type *T

	p := new(int)   // p, of type *int, points to an unnamed int variable
	fmt.Println(*p) // "0"
	*p = 2
	fmt.Println(*p) // "2"

	// A variable created with new is no different from an orginary local variable
	// except that there's no need to invent (and declare) a dummy variable

}

// the two newInt functions below have indentical behaviors

func newInt() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}

// new is a predeclared function, not a keyword
