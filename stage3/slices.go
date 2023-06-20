package main

import "fmt"

func main() {
	a := [5]int{0, 1, 2, 3, 4}
	s1 := a[0:3]
	s2 := a[2:5]
	fmt.Println(a, s1, s2)
	s1 = s1[0:4]
	fmt.Println(a, s1, s2)

	fmt.Println("len(s1): ", len(s1), "cap(s1): ", cap(s1)) // capacity is number of elements between the start of the slice and the end of the underlying array
	fmt.Println("len(s2): ", len(s2), "cap(s2): ", cap(s2))
	// append - add new values to slice even if at capacity
	// append returns a new slice but does not modify existing slice
	// you must reassign the value
	s2 = append(s2, 5)
	fmt.Println(a, s1, s2)
	// if you go over capacity, append will create a new underlying array
	s2[0] = 999
	fmt.Println(a, s1, s2)

	// better way to create lists of items in Go
	slice := []int{1, 2, 3} // underlying array already created for the slice
	slice = append(slice, 4, 5)
	slice = append(slice, 6, 7, 8)
	fmt.Println(slice)

	for i, v := range slice {
		fmt.Println("element: ", i, "value: ", v)
	}
}

func First(slice []int) (int, error) {
	if len(slice) > 0 {
		return slice[0], nil
	} else {
		return 0, fmt.Errorf("Slice is empty")
	}
}

func PrintEach(slice []string) {
	for _, v := range slice {
		fmt.Println(v)
	}
}
