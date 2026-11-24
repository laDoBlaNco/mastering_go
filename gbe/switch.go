//go:build ignore

package main

import (
	"fmt"
	"time"
)

// Switch statements express conditionals across many
// branches (3  or more)

func main() {
	// basic switch
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	fmt.Println()
	
	// I can use commas to separate multiple expressions in the same case statement
	// We can use the optional default as well
	switch time.Now().Weekday() {
	case time.Saturday,time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	fmt.Println()
	
	/* switch without an expression is an alternate way to express if/else
	logic. Here we also show how the case expression can be non-constants*/
	t:=time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	fmt.Println()
	
	/* A type switch comapares types instead of values. This is advanced Go and I think 
	I learned about this when studying generics. We can use this to discover the type of
	an interface value. Here the variable t will have the type corresponding to its
	clause */
	whatAmI := func(i any){
		switch t:=i.(type){
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n",t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
