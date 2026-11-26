//go:build ignore

package main

import "fmt"

// this fact function calls itself until it reaches the base case oF fact(0)
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1) // classic recursion example
}

func main() {
	fmt.Println(fact(7))
	
	/*
	   Anony functions can also be recursive, but this requires explicitly declaring a variable
	   with 'var'to store the function before its defined
	*/
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		// since fib was previously declared in main, Go can figure out which function
		// to call here
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))

}
