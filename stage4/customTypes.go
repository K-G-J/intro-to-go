package main

import "fmt"

type Minutes int
type Hours int
type Weight float64
type Title string
type Answer bool

func main() {
	minutes := Minutes(37) // use type conversion to convert from underlying type to custom
	hours := Hours(4)
	weight := Weight(945.7)
	name := Title("The Matrix")
	answer := Answer(true)
	fmt.Println(minutes, hours, weight, name, answer)
	minutes += 3
	fmt.Println(minutes)
	name += " Movie"
	fmt.Println(name)

	// can be compared to same custom type or to same underlying type
	// BUT you CANNOT compare custom types to custom types that have same underlying values
	if weight > Weight(907.3) {
		fmt.Println("Capacity exceeded")
	}
	if weight > 907.3 {
		fmt.Println("Capacity exceeded")
	}
}
