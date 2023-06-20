package main

import "fmt"

// you must use every variable you declare
// variables must begin with letters
// upper or lower case decides whether it is exportable
// cannot assign variables to different type without type conversion
// cannot to comparisons of different types
// variables are limited to block scopes and package-level
func main() {
	var a int
	a = 1
	var b, c int
	b, c = 2, 3
	var d = 4
	e, f := 5, 6
	fmt.Println(a, b, c, d, e, f)

	var wholeNumber int = 1
	var fractionalNumber float64 = 2.75
	var wholeNumber2 int = int(fractionalNumber)
	var fractionalNumber2 float64 = float64(wholeNumber)
	fmt.Println("wholeNumber2:", wholeNumber2)
	fmt.Println("fractionalNumber2:", fractionalNumber2)
	fmt.Println(float64(wholeNumber) + fractionalNumber)
	fmt.Println(float64(wholeNumber) < fractionalNumber)
}
