//Verifying pass by value

package main

import "fmt"

func main() {
	num := 10
	fmt.Println("num:\t", num, "\tAddress:\t", &num, "\t")

	increment(num)

	fmt.Println("num:\t", num, "\tAddress:\t", &num, "\t")
}

func increment(num int) {
	fmt.Println("num:\t", num, "\tAddress:\t", &num, "\t")
	num++
	fmt.Println("num:\t", num, "\tAddress:\t", &num, "\t")

}
