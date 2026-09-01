package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	years := 10.0
	expectedReturnRate := 5.5

	fmt.Print("Enter your investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)
	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, float64(years))
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
