package main

import "fmt"

func main() {
	var aValue float64 = 1.23
	var aPointer *float64 = &aValue
	fmt.Println("aPointer: ", aPointer)
	fmt.Println("aPointer: ", *aPointer)

	myNumber := 2.6
	halveNoPointer(myNumber)                      // Does nothing!
	fmt.Println("myNumber in 'main': ", myNumber) // Print 2.6

	halveWithPointer(&myNumber)
	fmt.Println("myNumber in 'main': ", myNumber) // 1.3 update value at address
}

func halveNoPointer(number float64) {
	number = number / 2
	fmt.Println("number in 'halve': ", number) // Prints 1.3, but change in only effective here!
}

func halveWithPointer(number *float64) {
	*number = *number / 2
	fmt.Println("*number in 'halve': ", *number) // 1.3
}
