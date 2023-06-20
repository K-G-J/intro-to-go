package main

import "fmt"

type Minutes int
type Hours int

func main() {
	minutes := Minutes(37)
	hours := Hours(4)
	fmt.Println(hours, " hours and ", minutes, " minutes")
}
