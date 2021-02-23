//Verifying pass by value

package main

import "fmt"

func main() {
	num := 10
	fmt.Println("num:\t", num, "\tAddress:\t", &num, "\t")
	//We passed the address here
	increment(&num)
	//Back to the stack, the value in the address has changed, so the value here is incremented.
	fmt.Println("num:\t", num, "\tAddress:\t", &num, "\t")
}

func increment(num *int) {
	// *num is how we access the value inside the address passed from the main
	fmt.Println("num:\t", *num, "\tAddress:\t", num, "\t &num:\t", &num)
	*num++
	fmt.Println("num:\t", *num, "\tAddress:\t", num, "\t &num:\t", &num)

}
