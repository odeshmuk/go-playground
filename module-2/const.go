package main

import "fmt"

func main() {
	// Untyped Constants.
	const ui = 12345    // kind: integer
	const uf = 3.141592 // kind: floating-point

	fmt.Println(ui)
	fmt.Println(uf)

	// Typed Constants still use the constant type system but their precision is restricted.
	const ti int = 12345        // type: int
	const tf float64 = 3.141592 // type: float64

	fmt.Println(ti)
	fmt.Println(tf)

	var answer = 3 * 0.333 // KindFloat(3) * KindFloat(0.333)
	fmt.Println(answer)

	var ans2 = 10. / 3 // Precision for div
	fmt.Println(ans2)

	const (
		A1 = iota // 0 : Start at 0
		B1 = iota // 1 : Increment by 1
		C1 = iota // 2 : Increment by 1
	)

	fmt.Println("1:", A1, B1, C1)

}
