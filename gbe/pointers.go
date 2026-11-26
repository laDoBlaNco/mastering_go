//go:build ignore

package main

/* Go supports pointers, allowing me to pass references to values and
records in our program

*/

import "fmt"

/*
	Here we'll see how pointers work in contrast to values with 2 functions:

zeroval and zeroptr. zeroval has an int parameter, so arguments will be passed by
value. zeroval will get a copy of ival distinct from the one in the calling
function
*/
func zeroval(ival int) {
	ival = 0
}

/*
	zeroptr in contrast has an *int parameter, meaning that it takes an int pointer. The

*iptr code in the function body then DEREFERENCES the pointer from its memory address to
the current value at that address. Assigning a value to a dereferenced pointer changes
the value at the referenced address.

so key to keep in mind that *T is the pointer type, but *variable is dereferencing
*/
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:",i)

	/* The &i syntax gives us the address of i, i.e. a pointer to i. So its basically
		*int -> pointer parameter in signature
		*i   -> dereferencing the ptr again in the definition
		&i   -> is the actual pointer (address) of the ARGUMENT
	*/
	zeroptr(&i)
	fmt.Println("zeroptr:",i)

	// Pointers can be printed too
	fmt.Println("pointer:",&i)

	// a new pointer to a value can be created with the builtin function new
	p:=new(42)
	fmt.Println("value at *p:",*p)
	zeroptr(p)
	fmt.Println("value at *p:",*p)
	
}
