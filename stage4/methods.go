package main

import (
	"fmt"
	"strings"
)

type Title string

// to create method on a type you give a special receiver to a function
func (t Title) FixCase() string {
	return strings.Title(string(t))
}

type Minutes int

// to modify the receiver, make a pointer
func (m *Minutes) Increment() {
	*m = (*m + 1) % 60
}

func main() {
	name := Title("the matrix")
	// use dot notation on the receiver
	fixed := name.FixCase()
	fmt.Println(fixed)

	minutes := Minutes(58)
	for i := 1; i <= 3; i++ {
		minutes.Increment()
		// (&minutes).Increment() -> Go automattically converts receiver to pointer for methods with pointer receiver
		fmt.Println(minutes)
	}
}
