package main

import "fmt"

func main() {
	if true {
		fmt.Println("true")
	}
	if false {
		fmt.Println("false")
	}
	if 1 < 2 {
		fmt.Println("1 < 2")
	}
	if 1 > 2 {
		fmt.Println("1 > 2")
	}
	if !true {
		fmt.Println("!true")
	}
	if !false {
		fmt.Println("!false")
	}
	if true && false {
		fmt.Println("true && false")
	}
	if true && true {
		fmt.Println("true && true")
	}
	if true || false {
		fmt.Println("true || false")
	}
	if false {
		fmt.Println("if")
	} else if true {
		fmt.Println("else if")
	} else {
		fmt.Println("else")
	}
}
