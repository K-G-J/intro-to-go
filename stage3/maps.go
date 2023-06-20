package main

import "fmt"

func main() {
	ages := map[string]float64{} // map of string to float64 that starts empty
	ages["Alice"] = 12
	ages["Bob"] = 9
	fmt.Println(ages)
	fmt.Println(ages["Alice"], ages["Bob"])

	agesPopulated := map[string]float64{"Alice": 12, "Bob": 9} // map of string to float64 that starts out with keys and values
	fmt.Println(agesPopulated)

	// for key, value in ages
	for name, age := range ages {
		fmt.Println(name, age)
	}
	// ignore keys
	for _, age := range ages {
		fmt.Println(age)
	}
	// ignore values
	for name := range ages {
		fmt.Println(name)
	}
}

func HalfPriceSale(prices map[string]float64) map[string]float64 {
	halfPrice := map[string]float64{}
	for product, price := range prices {
		halfPrice[product] = price / 2
	}
	return halfPrice
}
