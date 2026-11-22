//go:build ignore

/*
	In Go, variables are explicitly declared and used by the

compiler to e.g. check type-correctness of functon calls
*/
package main

import "fmt"

func main() {
	// var declares 1 or more variables
	var a = "initial"
	fmt.Println(a)

	// we can declare multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding initialization are
	// zero-valued. for example the zero value of an int is 0
	var e int
	fmt.Println(e)

	/* The := syntax is shorthand for declaring and intializing a variable,
	for var f string = "apple" for example.
	Only available inside functions (including main) */
	f := "apple"
	fmt.Println(f)
}
