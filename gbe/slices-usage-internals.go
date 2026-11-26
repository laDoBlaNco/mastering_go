//go:build ignore

package main

import (
	"os"
	"fmt"
	
	"regexp"
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
	s = make([]byte, 5)                 // when the cap ar is omitted, ti defaults to the specified length
	s = []byte{'k', 'e', 'v', 'i', 'n'} // reassigning actual values
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
	b2 := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(b2)
	fmt.Println(b2[1:4])

	// The start end indices of slice are optional as in other langs
	// This is also the syntax to create a slice given an array
	x := [3]string{"one", "two", "three"}
	sli := x[:] // a slice referencing the storage location of x
	fmt.Println(sli)
	fmt.Println()

	/*
		Slice Internals:

		A slice is just a descriptor of an array segment, which is why we refer to
		it as a window into an array. It consists in reality of a pointer to the array,
		the length of the segment, and its capacity. (image:slice-header.png)

		So a variable s slice would be a 3-word header with a pointer to an underlying
		array of bytes with 5 spaces allocated make([]byte,5) (image:slice-header-and-backing-array.png)

		The length is the number of elements that the slice refers to. The capacity is the
		number of elements in the underlying array (beginning at the element referred to by
		the slice pointer - since capacity for that slice would need to start there). The
		distinction between length and capacity will be made clear as I walk through the next
		few examples

		As I slice s, I can observe the changes in the slice data structure and their relation to
		the backing array: (image:byte-slice-backing-array)

		Slicing doesn't copy teh slice's data. It creates a new slice value (header) that points
		to the original array (reference). This makes slice operations as efficient as manipulating
		array indices. Therefore, modifying the elements (not the slice itself) of a re-slice
		modifies the elements of the original slice (multiple views into the same backing array):
	*/
	d := []byte{'r', 'o', 'a', 'd'}
	fmt.Println(d)
	e := d[2:] // a reslice of d --> []byte{'a','d'}
	fmt.Println(e)
	e[1] = 'm' // --> []byte{'a','m'}
	fmt.Println(e)
	fmt.Println(d) // --> original slice changed
	fmt.Println()

	// Earlier I sliced s to a length shorter than its capacity. But we can grow it to its capacity by slicing
	// it again. Using its capacity as the ending index (image: reslice-to-capacity.png)
	// s = s[:cap(s)]

	/*
		A slice can't though be grown beyond its capacity. Attempting to do this will cause a runtime panic, just
		as when indexing outside the bounds of a slice or array. Similarly, slices can't be re-sliced below zero
		to access earlier elements in the array. The window is the window..

		Growing Slices (the copy and append functions)

		So to increase the capacity of a slice I must create a NEW, larger slice and copy the contents of the original
		slice into it. This technique is how dynamic array implementations from other languages actually
		work behind the scenes. The next example doubles the capacity of s by making a new slice t,
		copying the contents of s into t, and then assigning the slice value t to s:
	*/
	fmt.Println(s, len(s), cap(s))

	t := make([]byte, len(s), (cap(s)+1)*2) // +1 is just in case cap(s) == 0
	copy(t, s)                              // then we copy in t <- s
	s = t
	fmt.Println(s, len(s), cap(s))
	fmt.Println()

	/*
		The above was first shown in a loop (naive method), but this is made easier by the built-in
		copy function. As the name suggests, copy copies daa from a source slice to a destinatino slice. It
		returns the number of elements copied, if I need them. Its a little unintuitive because the signature
		is a bit backwards: func copy(dst,src []T) int  so dest is the first arg and src the second of same type
		of course

		The copy function supports copying betweeen slices of different lengths as well (it will copy only up to
		the smaller number of elements). In addition, copy can handle source and destination slices that share the
		same backing array, handling overlapping slices correctly. I think I lose that trying to do a manual loop

		A common operation is to append data to the end of a slice. This function (AppendByte) appends byte elements to a slice
		of bytes, growing the slice if necessary, and returns the updated slice value:
	*/
	fmt.Println()

	p := []byte{2, 3, 5}
	fmt.Println(p, len(p), cap(p))
	p = AppendByte(p, 7, 11, 13)
	fmt.Println(p, len(p), cap(p))
	fmt.Println()

	/*
		Functions liek AppendByte are useful because they offer complete control over the way a slice
		is grown. Depending on the characteristics of the program, it may be desirable to allocate in
		smaller or larger chunks, or to put a ceiling on the size of the reallocation.

		But most programs don't need complete control, so Go provides a built-in append function that's
		good for most purposes; its signature is the following:

		func append(s []T, x ...T) []T

		The built-in function appends the elements x to the end of the slice s, and grows the slice if
		a greater capacity is needed. REMEMBER growing the slice is a new allocation with a different
		memory address
	*/
	aSlice := make([]int, 1)
	fmt.Println(aSlice, len(aSlice), cap(aSlice))
	aSlice = append(aSlice, 1, 2, 3) // basically same formula (+1 and *2)
	fmt.Println(aSlice, len(aSlice), cap(aSlice))
	fmt.Println()

	// To append one slice to another, we use ... to expand or spread which is the counterpart to our
	// variadic function parameter in our AppendByte function '...byte'
	// the second argument to a list of arguments
	a3:=[]string{"John","Paul"}
	b3:=[]string{"George","Ringo","Pete"}
	a3 = append(a3,b3...) // same as - append(a3,b3[0],b3[1],b3[2])
	fmt.Println(a3,len(a3),cap(a3))
	fmt.Println()
	
	// Since the zero value of a slice (nil) acts like a zero-length slice, I can declare a slice
	// variable and then append to it in a loop:
	filter:=func(s []int, fn func(int)bool)[]int{
		var p []int // == nil zero-default
		fmt.Println(p) // logging p before the loop
		
		for _,v:=range s{
			if fn(v){
				p=append(p,v) // appending to p even though its nil
			}
		}
		return p
	}
	fs:=filter([]int{1,2,3,4,5,6,7,8,9,10},func(n int)bool{
		return n%2==0
	})
	fmt.Println(fs)
	fmt.Println()
	
	/* 
		Possible "Gotcha"

		As mentioned earlier, re-slicing a slice doesn't make a copy of the  backing array. The
		full array will be kept in memory until it is no longer referenced. Occasionally, this can
		cause the program to hold all the data in memory when only a small piece of it is needed.

		for example, below in FindDigits, we load a file into memory and search it for the first
		group of consecutive numeric digits, returning them as a new slice

		That code behaves as advertised, but the returned []byte points into a backing array containing
		the entire file. Since the slice references teh original array, as long as the slice is kept
		around the gc can't release the array; the few useful bytes of the file keep the entire contents
		in memory.

		To fix this problem I could just copy the interesting data to a new slice before returning 
		it (CopyDigits). A more concise function could also be used with append.

	*/
}
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string)[]byte{
	b, _ := os.ReadFile(filename)
	return digitRegexp.Find(b)
}

func CopyDigits(filename string)[]byte{
	b,_:=os.ReadFile(filename)
	b = digitRegexp.Find(b)
	c:=make([]byte,len(b))
	copy(c,b)
	return c
}