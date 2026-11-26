//go:build ignore

package main

import (
	"bytes"
	"fmt"
)

/*
Arrays, slices (and strings): The mechanics of append

One of the most common features of precedureal programming languages is the concept of
an array. Arrays seem like simple things but there are many questions that must be answered
when adding them to a language, such as:

	- fixed-size or variable-size?
	- is the size part of the type?
	- what do multidimensional arrays look like?
	- does the empty array have meaning?

The answers to these questions affect whether arrays are just a feature of the language
or a core part of its design.

In the early development of Go, it took about a year to decide the answers to these
questions before the design felt right. The key step was the itnroduction of 'slices',
which built on top of fixed-sized arrays to give a flexible, extensible data structure.
To this day, however, programmers new to Go often stumble over the way slices work,
perhaps becasue experience from other languages has colored their thinking.

In this post we'll atempt to clear up the confusion. We'll do so by building up the pieces
to explain how the append built-in function works, and why it works the way it does.

Arrays

Arrays are an important building block in Go, but like the foundation of a building they
are often hidden below more visible components. We must talk about them briefly before we move
on to the more interesting, powerful, and prominent idea if slices.

Arrays are not often seen in Go programs because the size of an array is part of its type,
which limits its expressive power.
*/

func main() {

	// The declaration of an array
	var buffer [256]byte

	/*
		declares the variable buffer, whic holds 256 bytes. The type of buffer includes its
		size, [256]byte. An array with 512 bytes would be of the distinct type [512]byte.

		The data associated with an array is just that: an array of elements. Schematically,
		our buffer looks like this in memory

			buffer: byte byte byte byte ... 256 times ... byte byte byte

		That is it, the variable holds 256 bytes of data and nothing else. We can access its
		dlements with the familiar indexing syntax buffer[0], buffer[1], etc through buffer[255].
		(The index range 0 through 255 covers 256 elements). Atempting to index buffer with
		a value outside this range will crash the program in 'panic'

		There is a built-in function called 'len' that rturns the number of elements of an
		array or slice and alos of the few other data types. For arrays, it's obvious what
		len returns. In our example, len(buffer) returns 256

		Arrays have their place, they are a good reprsentation of transformation matrix for
		instance, but their most common purpose in Go is to hold storage for a slice.

		Slices: the slice header
		Slices are whre the action is, but to use them, well one must understand exactly what
		they are and what they do.

		A slice is a data structure describing a contiuous section of an array sxtored separately
		from the slice variable itself. A slice IS NOT AN ARRAY.  A slice DESCRIBES a PIECE OF
		AN ARRAY.

		Given our buffer array variable from the prev ious section, we could create a slice
		that describes elements 100 through 150 (to be precise, 100 through 149, inclusive)
		by slicing the array
	*/
	// var slice []byte = buffer[100:150]
	// fmt.Println(slice)

	/*
		In that snippet we used the full variable declaration to be explicit. The variable
		slice has type []byte, pronounced "slice of bytes", and is initialized from the
		array, called buffer, by slicing elements 100 (inclusive) thorugh 150 (exclusive).
		The more Go idomatic syntax would drop the type, which is set by the initializing
		expression: var slice = buffer[100:150]. And inside a function, the short declaration
		slice := buffer[100:150]

		What exactly is this slice variable? Is not quite the full story, but for now think
		of it as a little data structure with two elements: a length and a pointer to an element
		of an array. We can think of it as being built like this behind the scenes
	*/
	/*
		type sliceHeader struct {
			Length        int
			ZerothElement *byte
		}

		slice := sliceHeader{
			Length:        50,
			ZerothElement: &buffer[100],
		}
	*/
	/* Of course this is just an illustration. Despite what this snippet says, that
	sliceHeader struct is not visible to the programmer, and the type of the element
	pointer depends on the type of the elements, but this gives the general mental model
	of the mechanics.

	So far we've used a slice operation on an array, but we can also slice a slice like
	the following
	*/
	// slice2 := slice[5:10]
	// fmt.Println(slice2)

	/*
		Just as before, this operation creates a new slice, in thi scase with elements 5
		through 9 (inclusive) of the original slice, which means elements 105 through 109
		of the original backing array. The underlying sliceHeader struct for the slice2 variable
		would looke like this:

			slice2 := sliceHeader{
				Length: 5,
				ZerothElement: &buffer[105],
			}

		Notice that this header still points to the same backing array, stored in the
		buffer variable.

		We can then also reslice, o sea, slice a slice and store that result back in the
		original slice structure. 🤯
	*/
	// slice = slice[5:10]

	/* The sliceHeader structure for the slice variable looks just like it did for the
	slice2 variable. We'll see reslicing used often, for example to truncate a slice.
	This statement drops the first and last elements of our slice
	*/
	// slice = slice[1:len(slice)-1]

	/*
		Excercise:
		The slice header would now look like
		slice = sliceHeader{
			Length: 4
			ZerothElement: &buffer[106]
		}

		I'll often here in the wild Go programmers talk about 'slice header'because that
		really is what's stored in a slice variable. For instance, when we call a function
		that takes a slice as an argument, such as bytes.IndexRune, that header is what gets
		passed to the function.

			slashPos := bytes.IndexRune(slice, '/')

		The slice argument that is passed to the IndexRune function is, in fact a slice
		header.

		There's one more data item in the slice header, which we'll talk about below, but
		first let's see what the existence of the slice header means when we program in Go
		with slices.

		Passing slices to functions

		Its important to understand that even though a slice contains a pointer, it is itself
		a value. Under the covers, it is a struct value holding a pointer and a length. It is
		not a pointer to a struct. So we true are still passing values.

		This matters

		When we called IndexRune in the previous example, it was passed a COPY of the slice
		header. That behavior has important ramifications.

		consider the following function below AddOneToEachElement
	*/

	/* It does just what its name implies, iterating over the indices of a slice (using
	a for range loop), incrementing its elements
	*/

	slice := buffer[10:20]
	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}
	fmt.Println("before", slice)
	AddOneToEachElement(slice)
	fmt.Println("after", slice)

	/* Even though the slice header is passed by value (copy), the header includes a pointer
	to elements of an array, so both the original slice header and the copy of the header
	passed to the function describe the same array. Therefore, when the function returns,
	the modified elements can be seen through the original slice (window) variable

	The argument really is a copy as we can see from the example below:*/
	fmt.Println("Before: len(slice) =", len(slice))
	newSlice := SubtractOneFromLength(slice)             // we run the function and get a different result returned
	fmt.Println("After: len(slice) =", len(slice))       // our original window into buffer
	fmt.Println("After: len(newSlice) =", len(newSlice)) // our new window into buffer (copy)

	/*
		Here we see that the contents of slice argument can be modified by a function, but its header
		cannot. The length stored in the slice variable is not modified by the call to the function,
		since that function is passed by copy of the slice header, not the original. Thus if we
		want to write a function that modifies the header, we must return it as a result parameter,
		just as we have done here. The slice variable is unchanged but the returned value (copy)
		has the new length, whic is then stored in newSlice

		Pointers to slices: Method receivers

		Another way to have a function modify the slice header is to pass a pointer directly to
		it. Here's a variant of our previous example that does just this:
	*/
	fmt.Println("Before: len(slice) =", len(slice))
	PtrSubtractOneFromLength(&slice)
	fmt.Println("After: len(slice) =", len(slice))

	/*
		It seems clumsy in that example, espcially dealing with the extra level of indirection
		(a temporary variable helps), but there's one common case where we see pointers to slices.
		It is idiomatic to use a pointer reciever for a method that modifies a slice. So this makes
		sense as it would be the difference between sort() and toSorted() or reverse() and toReversed()
		in Javascript. sort and reverse mutate the actual array while the others return a new one.
		So in Go to have a method actually changing the receiving type it needs a pointer to that
		receiver. 🤓🤓🤓

		let's say we wanted to have a method on a slice that truncates it at the final slash. We
		could the following:
	*/
	pathName := path("/usr/bin/tso") // conversion from string to path
	pathName.TruncateAtFinalSlash()
	fmt.Printf("%s\n", pathName)

	/*
		When I run this example I see that it works properly, updating the slice in the caller.
		As an excercise I passed in a value instead of a pointer to our receiver and nothing was
		changed. I would then need to RETURN the new path

		On the other hand, if I wanted to write a method for path that upper-cases the ASCII
		letters in the path (parochially ignoring non-English names), the method could be a value
		value since the value receiver will still point to the same underlying array.
	*/
	pathName.ToUpper()
	fmt.Printf("%s\n", pathName)

	/*
		Im not sure why this works 🤔🤔.
		Here the ToUpper method uses two variables in the for range construct to capture the
		index and slice element. This form of loop avoids writing p[i] multiple times in the body

		I understand why it works, we are just muting a slice which is referencing the same backing
		array, I get that. But then why did we have to do the pointer reciever on the other? What's
		the difference?

		NOTE - Got with Gemini and now I get it. In ToUpper we are mutating a value in the slice
		which is mutating the backing array as normal and as I put above. With the Truncate function
		we aren't mutating the backing array, we are RESLICING, meaning we are only changing the
		actual slice header. So if we don't use a pointer, we just change the length on our passed
		in copy but the original doesn't change. Using a pointer reciever we are changing the
		original window so we can see a different part of our backing array.

		In the end though, its better to use one semantic value or pointer and stick to it. And
		idiomatic Go uses the semi-fp way of simply sticking with value semantics and returning
		the new value.

		Capacity:
		Now let's look at the function Extend below. It extends its argument slice of ints
		by one element. We need to return the the modified slice since we aren't using pointers
		and only modifying the slice header in that step. When we run it, it seems to work, until
		it hits the capacity wall of 10. runtime error: slice bounds out of range [:11] with capacity 10

		See how the slice grows...until it doesn't??

		So now its time to dig into that 3rd part of the slice header, its CAPACITY. Besides
		the array pointer and length, the slice header also stores its capacity:

			type sliceHeader struct{
				Length int
				Capacity int
				ZerothElemenet *byte
			}

		The capacity field records how much space the underlying array actually has; it is the
		maximum value the Length can reach. Trying to grow the slice beyond its capacity will
		step beyond the limits of the backing array and will trigger a panic

		After our example slice is created with slice2 := iBuffer[0:0], its header is:

			slice2 := sliceHeader{
				Length: 0,
				Capacity: 10,
				ZerothElement: &iBuffer[0],
			}

		The capacity field is equal to the length of the underlying array, minus the index
		in the array of the first element of the slice (in this case its full capacity since we 
		are at zero) If our slice started at 2 then capacity would only be 7. If I want to inquire
		what the capacity is for a slice, I have the built-in function 'cap':
	*/
	var iBuffer [10]int
	slice2 := iBuffer[0:0]
	for i := 0; i < 20; i++ {
		if cap(slice2) == len(slice2){
			fmt.Println("slice is full")
			break
		}
		slice2 = Extend(slice2, i)
		fmt.Println(slice2)
	}

	/* 
	Make: 

	
	*/
}

func Extend(slice []int, element int) []int {
	n := len(slice)
	slice = slice[0 : n+1] // again we are reslicing but in the other direction - slice header manipulation
	slice[n] = element     // here we are manipulating the underlying array
	return slice           // and we return the new slice (window/view) of that array
}

type path []byte

func (p path) ToUpper() {
	for i, b := range p {
		if 'a' <= b && b <= 'z' {
			p[i] = b + 'A' - 'a'
		}
	}
}

// NOTE this is what I read about we typically do methods on structs as we would objects in
// other languages, but we can put a method on any type. here we put a method on a new type of
// []byte named path. This seems similar to maybe Number.prototype.path = function(){} in JS
// then run 1..path() and it'll work. We are taught that methods only go with objects OOP,
// but we can put a method or associated function on any type at all
// this method receiver pattern is similar to using 'this' or 'self' in oop from what I can tell
func (p *path) TruncateAtFinalSlash() {
	i := bytes.LastIndex(*p, []byte("/"))
	if i >= 0 {
		*p = (*p)[0:i] // ()s to index from a pointer
	}
}

func PtrSubtractOneFromLength(slicePtr *[]byte) {
	slice := *slicePtr
	*slicePtr = slice[0 : len(slice)-1]
}

func SubtractOneFromLength(slice []byte) []byte {
	slice = slice[0 : len(slice)-1]
	return slice
}

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}
