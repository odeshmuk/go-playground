package main

import "fmt"

func main() {
	
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n\n", d, d)

	// Using the short variable declaration operator, we can define and initialize at the same time.
	// aa := 10
	var aa=10
	bb := "hello" // 1st word points to a array of characters, 2nd word is 5 bytes
	cc := 3.14159
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n\n", dd, dd)

	// ---------------------
	// Conversion vs casting
	// ---------------------

	// Go doesn't have casting, but conversion.
	// Instead of telling a compiler to pretend to have some more bytes, we have to allocate more
	// memory.

	// Specify type and perform a conversion.
	aaa := int32(10)

	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)
}