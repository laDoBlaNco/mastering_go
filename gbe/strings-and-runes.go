//go:build ignore

package main

import "fmt"

/*
A Go string is az read-only slice of bytes. The language and teh standard library treat
strings specially - as containers of text encoded in UTF-8. In other languages, strings
are made of "characters". In Go, the concept of a character is called a "rune" - it's an
integer that represents a Unicode code point.
*/

func main() {
	/*
		s is a string assigned a literal value representing the word "hello" in the Thai
		language. Go string literals are UTF-8 encoded text
	*/
	const s = "สวัสดี"

	/*
		Since strings are equivalent to []byte, this will produce the length of the raw
		bytes stored within
	*/
	fmt.Println("len:", len(s))

}
