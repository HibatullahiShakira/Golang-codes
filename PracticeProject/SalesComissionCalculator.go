package main

import "fmt"

func main() {
	const weeklyEarnings float64 = 200
	fmt.Println("How much worth of merchandise did the salesperson sell? ")
	var merchandiseWorth float64
	fmt.Scanf("%f", &merchandiseWorth)
	percentageOfMerchandiseEarned := merchandiseWorth * 0.09
	fmt.Println(percentageOfMerchandiseEarned)
	totalWeeklyAmountEarned := percentageOfMerchandiseEarned + weeklyEarnings
	fmt.Printf("The total amount earned weekly by the salesperson is %.2fn", totalWeeklyAmountEarned)
}
