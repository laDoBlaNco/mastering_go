//go:build ignore

package main

/* 
Strings, bytes, runes and Characters in Go

At first, strings might seem like too simply a topic for a blog post, but to use them
well requires understanding not only how they work, but also the difference between a
byte, a character, and a rune, the difference between Unicode and UTF-8, the difference
between a string and a string literal, and other even more subtle distinctions.

One way to approach this topic is to think of it sas an answer to the frequently asked
question, "When I index a Go string at position n, why don't I get the nth character?"
As we'll see, this question leads us to many details about how text works in the modern
world.

What is a string?
Starting with the basics.

In Go, a string is in effect a read-only slice of bytes. 
*/