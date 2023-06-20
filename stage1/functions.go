package main

import "fmt"

func main() {
	myFunction()
	fmt.Println(square(3))
	fmt.Println(add(2, 4))
	fmt.Println(subtract(10, 3))
}

// function name must start with letter
// uppercase first letter is exported
func myFunction() {
	fmt.Println("Running myFunction")
}

func square(number int) int {
	return number * number
}

func add(a float64, b float64) (sum float64) {
	return a + b
}

// named return
func subtract(a, b float64) (difference float64) {
	difference = a - b
	return
}
