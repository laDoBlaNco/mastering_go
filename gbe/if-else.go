//go:build ignore

package main

import "fmt"

// Branching with if and else in Go is straight-forward

func main() {
	// Here's a basic example
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	fmt.Println()

	// I can even have an if statement without an else
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
	fmt.Println()

	// Logical operators like && and || are often useful in conditions
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are event")
	}
	fmt.Println()

	// A statement can precede conditionals; any variables declared in this statement
	// are available in the current and all subsequent branches
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
/* Note that I don't need ()s around conditions in Go, but braces are required

Also, there is no ternary if in Go, so I'll need to use the full if statement even for
basic conditions.*/
