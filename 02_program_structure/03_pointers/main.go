package main

import "fmt"

func main() {
	// A variable is a piece of storage containing a value.
	// A pointer is the address of variable. A pointer is thus the location
	// at which a value is stored
	// With a pointer, we can read or update the value of a variable inderectly,
	// without using or even knowing the name of the variable.

	// If a variable is declared var x int, the expression &x ("address of x") yiels
	// a pointer to an integer variable.
	// That is, a value of type *int, which is a pronounced "pointer to int"
	// If this value is called p, we say "p points to x"
	// The expression *p yields the value of that variable

	x := 1
	p := &x         // p, of type *int, points to x
	fmt.Println(*p) // "1"
	*p = 2          // equivalent to x = 2
	fmt.Println(x)  // 2

	// Pointers are comparable, two pointers are equal if and only if they point to the
	// same variable or both are nil
	// var x, y int
	// fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"

}
