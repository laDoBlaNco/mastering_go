//go:build ignore

package main

import (
	"fmt"
	"slices"
)

// Slices are an important data type in Go, giving a more powerful interface to sequences
// than arrays. They are windows into an underlying array.

func main() {
	/* Unlike arrays, slices are typed only byt the elements they contain (not
	the number of elements). An uninitialized slice equals to nil and has length 0
	*/
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)
	fmt.Println()

	/*
		To create a slice with non-zero length, I use the builtin 'make'. Here we will
		make a slice of strings of length 3 (initially with the zero-default). By default
		a new slice's capacity is equal to its length; if we know the slice is going to
		grow ahead of time, it's possible to pass a capacity explicitly as an additional
		parameter to 'make'
	*/
	s = make([]string, 3)
	fmt.Println("emp;", s, "len:", len(s), "cap:", cap(s))
	fmt.Println()

	// we can set and get just like arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println()

	// len returns  the length of the slice as expected
	fmt.Println("len:", len(s))

	/*
		In addition tyo these basic operations, slices support several more that make them
		richer than arrays. One is the builtin 'append', which returns a slice containing one
		or more new values. Note that we need to accept a return value from append as we MAY
		get a new slice value.
	*/
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	fmt.Println()

	/*
		Slices can also be copied. Here we create an empty slice c of the same length as
		s and copy into c from s
	*/
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)
	fmt.Println()

	/*
		Slices support a "slice" operator with the syntax slice[low:high]. For example,
		this gets a slice of the elements s[2], s[3], and s[4]
	*/
	l := s[2:5]
	fmt.Println("sl1:", l)
	fmt.Println()

	// This slices up to (but excluding) s[5]
	l = s[:5]
	fmt.Println("sl2:", l)
	fmt.Println()

	// and this slices up from (and including) s[2]
	l = s[2:]
	fmt.Println("sl3:", l)
	fmt.Println()

	// we can also declare and initialize a variable for slice in a single line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	fmt.Println()

	// The slices package contains a number of useful utility functions for slices
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}
	fmt.Println()
	
	/* 
		Slices can be composed into multi-dimensional data structures. The length
		of the inner slices can vary, unlike with multi-dimensional arrays
	*/
	twoD := make([][]int,3)
	for i := range 3{
		innerLen:=i+1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen{
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	
}
