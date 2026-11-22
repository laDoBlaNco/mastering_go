//go:build ignore

package main

import "fmt"

/*
for is Go's ONLY looping construct. Here are some basic types of for loops
*/

func main() {
	// The most basic type, with a single condition - like while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}
	fmt.Println()
	// A classic initial/condition/after for loop
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
	fmt.Println()
	// Another way of accomplishing the basic "do this N times"is
	// range over an integer iteration
	for i := range 3 {
		fmt.Println("range", i)
	}
	fmt.Println()
	/*
		for without a condition will loop repeatedly until I break out of the loop
		or return from the enclosing function
	*/
	for {
		fmt.Println("loop")
		break
	}
	fmt.Println()
	// we can also continue to the next iteration of the loop
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)

	}

	// I'll see some other 'for' forms later wiwth range statements, channels,
	// and other data structures

}
