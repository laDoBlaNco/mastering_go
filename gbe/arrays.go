//go:build ignore

package main

import "fmt"

/* In Go, an array is a numbered sequence of elements of a specific length.
In typical Go code, slices are much more common; arrays are useful in some special
scenarios */

func main() {
	/* Here we create an array 'a' that will hold exactly 5 ints. The type of elements
	and length are both part of the array's type name. By default an array is zero-valued
	wwhic for ints means 0s */
	var a [5]int
	fmt.Println(a)
	fmt.Println()
	
	/* We can set a value at an index using the array[index] = value syntax
	and get a value wiwth array[index] */
	a[4] = 100
	fmt.Println("set:",a)
	fmt.Println("get:",a[4])
	fmt.Println()
	
	/* The builtin len returns the length of an array */
	fmt.Println("len:",len(a))
	fmt.Println()
	
	/* use this syntax to declare and initialize an array in one line. */
	b:=[5]int{1,2,3,4,5}
	fmt.Println("dcl:",b)
	fmt.Println()
	
	/* We an also have the compiler count the number of elements for us. This is still
	a static array */
	b = [...]int{1,2,3,4,5,}
	fmt.Println("dcl:",b)
	fmt.Println()
	
	/* If we specify the index with :, the elements in between will be zeroed */
	b = [...]int{100,3:400,500}
	fmt.Println("idx:",b)
	fmt.Println()
	
	/* Array types are one-dimensional, but we can compose types to build multi-dimensional
	data structures */
	
	var twoD [2][3]int
	for i:=range 2{
		for j:=range 3{
			twoD[i][j] = i+j
		}
	}
	fmt.Println("2d: ",twoD)

	// We can create an initialize multi-dimensional arrays at once too
	twoD = [2][3]int{
		{1,2,3},
		{1,2,3},
	}
	fmt.Println("2d: ",twoD)

}
