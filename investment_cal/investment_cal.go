package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5

	var investmentAmount, years float64
	var expectedReturnRate float64

	fmt.Print("provide investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println("futureRealValue:", futureRealValue)

}
