package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	squareRoot, err := squareRoot(9)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(squareRoot)
	// assigning blank to unused return values
	fileInfo, error := os.Stat("existent.txt")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(fileInfo.Size())
	}

	fileInfo, error = os.Stat("nonexistent.txt")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(fileInfo.Size())
	}
}

func squareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("can't take square root of negative number")
	}
	return math.Sqrt(x), nil
}
