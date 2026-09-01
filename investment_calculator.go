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

	var revenue, expenses, taxRate float64

	fmt.Print("\nSelect an option: \n 1. Investment Calculator\n 2. Profit Calculator\n Enter your option: ")
	var choice string
	fmt.Scan(&choice)

	switch choice {
	case "1":
		fmt.Print("\nEnter your investment amount: ")
		fmt.Scan(&investmentAmount)
		fmt.Print("\nEnter the number of years: ")
		fmt.Scan(&years)
		fmt.Print("\nEnter the expected return rate: ")
		fmt.Scan(&expectedReturnRate)

		futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, float64(years))
		futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

		fmt.Println("Future value: ", fmt.Sprintf("%.2f", futureValue))
		fmt.Println("Future real value: ", fmt.Sprintf("%.2f", futureRealValue))
	case "2":
		fmt.Print("\nEnter revenue: ")
		fmt.Scan(&revenue)
		fmt.Print("\nEnter expenses: ")
		fmt.Scan(&expenses)
		fmt.Print("\nEnter tax rate: ")
		fmt.Scan(&taxRate)

		profit(revenue, expenses, taxRate)
	default:
		fmt.Println("Invalid option")
	}
}

func profit(revenue float64, expenses float64, taxRate float64) {
	profit := (revenue - expenses)
	tax := profit * (taxRate / 100)
	ratio := (profit / revenue) * 100
	fmt.Println("Earnings before tax: ", fmt.Sprintf("%.2f", profit))
	fmt.Println("Tax: ", fmt.Sprintf("%.2f", tax))
	fmt.Println("Net earnings: ", fmt.Sprintf("%.2f", profit-tax))
	fmt.Println("Ratio: ", fmt.Sprintf("%.2f", ratio), "%")
}
