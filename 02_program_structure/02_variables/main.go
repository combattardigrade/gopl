package main

func main() {
	// var i, j, k int                 // int, int, int
	// var b, f, s = true, 2.3, "four" // bool, float64, string

	// A set of variables can also be initialized by calling a function that
	// multiple value
	// var f, err = os.Open(name) // os.Open returns a file and an error

	// Short Variable Declarations
	// local variables
	// anim := gif.GIF{LoopCount: nframes}
	// freq := rand.Float64() * 3.0
	// t := 0.0
	// i := 100

	// multi-variable declaration
	// i, j := 0, 1
	// := is a declaration, whereas = is an assigment.
	// A multi-variable declaration should not be confused with tuple assignment
	// in which each variable on the left-hand side is assigned the corresponding value
	// from the right-hand side
	// example:
	// i, j = j, i // swap value of i and j

	// A short variable declaration not necessarily declare
	// all the variables on its left-hand side. If some of them were already declared
	// in the same lexical block, then the short variable declaration acts like an
	// assignment to those variables
	// in, err := os.Open(infile)
	// ...
	// out, err := os.Create(outfile)

	// A short variable declaration acts like an a assigment only to variables that
	// were already declared in the same lexical block; declarations in an outer
	// block are ignored.
}
