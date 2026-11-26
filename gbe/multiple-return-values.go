//go:build ignore

package main

import "fmt"

// (int,int) in this signature shows that the function returns 2 ints
func vals()(int,int){
	return 3,7
}

func main(){

	/* 
	Go has built-in support for multiple return values. This feature is used often in idiomatic
	Go, for example to return both result and error values from a function.
	*/
	
	// here we have 2 different return values from the call with multiple assignment
	a,b:=vals()
	fmt.Println(a)
	fmt.Println(b)
	
	// here we only want 1 of the return values so we use '_'
	_,c:=vals()
	fmt.Println(c)
	


}