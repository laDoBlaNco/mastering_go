//go:build ignore

package main

import (
	"fmt"
)

func main() {
	/*
		Go's slice stype providfes a convenient and efficient means of working with sequences
		of typed data. Slices are very simliar  to arrays in other langs, but have some unusual
		and special properties.

		Arrays
		The slice type is an abstrction built on top of Go's array type, so to understand
		slices I must first understand arrays.

		An array type definition specifies a length and an element type. For example, the
		type [4]int represents an array of four integers. An array's size is fixed; its
		length IS PART OF ITS TYPE ([4]int and [5]int are different types). Arrays can  be
		indexed in the usual way, so the expression s[n] accesses the nth element starting
		from zero as expected.
	*/
	var a [4]int
	a[0] = 1
	i := a[0]
	fmt.Println(i)

	/*
		Arrays don't need to be initialized explicitly; teh zero value default of an array
		is a ready-to-use array whose elements are themselves zeroed:
	*/
	fmt.Println(a[2] == 0)

	// The in-memory representation of [4]int is just four integer values laide out

	/*
		Goe's arrays are values. An array variable denotes the entire array; it IS NOT A
		POINTER TO THE FIRST ARRAY ELEMENT as in C. This means that when I assign or pass
		around an array value I'm making a copy of its contents. (To avoid the copy I
		would need to pass a pointer to the array, but that that's a pointer to an array
		and not an array). One way to think about arrays is as a sort of struct but  with
		indexed rather than named fields: a fixed size composite value
	*/

	// an array literal can be specified like this:
	b := [2]string{"Penn", "Teller"}
	// or I can have the compiler coiunt the array elements for me. Still typed and
	// known at compile time
	b = [...]string{"Penn", "Teller"}

	fmt.Println(b)

	/*
		Slices

		Arrays hve their place in Go, but they're a bit inflexible, so I won't see them
		too often in Go code. Slices, though, are everywhere. They build on arrays to
		provide great power and convenience.

		The type specification ofr a slice is []T, wwhere T is the type of the elements
		of the slice. Unlike an array type, a slice type has no specified length.

		A slice literal is declared just like an array literal, except I leave out the
		element count:
	*/
	letters := []string{"a", "b", "c", "d", "e"}
	fmt.Println(letters)

	/*
		A slice can be created with the built-in function called make, which has
		the signature: 
		func make([]T, len, cap) []T

		where T stands for  the element type of the slice to be created. The make function
		takes a type, a length, and optional capacity. When called, make allocates an array
		and returns a slice that refers to that array
	*/
	var s []byte
	s = make([]byte,5) // when the cap ar is omitted, ti defaults to the specified length
	fmt.Println(s)

	// the length and capacity can be inspected with the built-ins
	fmt.Println(len(s))
	fmt.Println(cap(s))

	/* 
	The next sections we'll discuss the relationship between length and capacity

	The zero value of a slice is nil. The len and cap functions will both return 0 for
	a nil slice

	A slice can also be formed by "slicing" an existing slice or array. Slicing  is
	done by specifying a half-open range with two indices separated by a :. For example
	b[1:4] creates a slice including elements 1-3 fo b (and teh indices of the resulting
	slice will be 0-2 as expected since its a different slice)

	*/
	b2:=[]byte{'g','o','l','a','n','g'}
	fmt.Println(b2)
	fmt.Println(b2[1:4])

	// The start end indices of slice are optional as in other langs
	// This is also the syntax to create a slice given an array
	x:=[3]string{"one","two","three"}
	sli:=x[:] // a slice referencing the storage location of x
	fmt.Println(sli)

	/* 
	Slice Internals:

	
	*/
	
		
	
	
}
