package main

import "fmt"

func main() {

	// productNames := [4]string{"dinesh", "ramesh", "suresh", "mahesh"}
	// //create array of 4 elements
	// prices := [4]float64{10.99, 9.99, 45.99, 10.0}

	// fmt.Println(prices)
	// fmt.Println(productNames)

	// //slices are dynamic arrays
	// featuredPrices := prices[1:3]
	// fmt.Println(featuredPrices)

	// fmt.Println(len(featuredPrices), cap(featuredPrices))

	// featuredPrices[0] = 1.99
	// highlightedPrices := featuredPrices[:1]
	// fmt.Println(highlightedPrices)
	// //fmt.Println(prices)
	// fmt.Println(len(highlightedPrices), cap(highlightedPrices))
	// fmt.Printf("\n")
	// highlightedPrices = highlightedPrices[:3]
	// fmt.Println(highlightedPrices)
	// fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	//Practice arrays and slices

	hobbies := [3]string{"reading", "writing", "coding"}
	fmt.Println(hobbies)

	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	mainHoobies := hobbies[:2]
	fmt.Println(mainHoobies)

	fmt.Println(cap(mainHoobies))
	mainHoobies = mainHoobies[1:3] //extend the slice to 3 elements * important in slice (underlying slice)
	fmt.Println(mainHoobies)

	//dynamic arrays (slices)
	coursegoals := []string{"learn go", "learn js", "learn python"}
	fmt.Println(coursegoals)

	coursegoals[1] = "learn java"
	fmt.Println(coursegoals)

	coursegoals = append(coursegoals, "learn c++")
	fmt.Println(coursegoals)

}
