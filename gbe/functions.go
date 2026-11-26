//go:build ignore

package main

import "fmt"

func main() {
	// Functions are central in Go. They are 'first-class citizens'
	res := plus(1, 2)
	fmt.Println("1 + 2 =", res)

	res = plusPlus(1,2,3)
	fmt.Println("1 + 2 + 3 =",res)
	
}

// Named functions must be at package level. Anony funcs can be inside other funcs
// This function taks two ints and returns their sum as an int
func plus(a int, b int) int {
	return a + b // Go requires explicit returns, it won't automatically return the value
	// of the last expression
}

/*
When I have multiple consecutive parameters of the same type, I can omit the type name for
the like-typed params up to the final param that declares the type
*/
func plusPlus(a, b, c int) int {
	return a + b + c
}
