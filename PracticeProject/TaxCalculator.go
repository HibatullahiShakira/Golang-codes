package main

import "fmt"

func main() {
	fmt.Println("How much is your monthly salary")
	var monthlySalary float64
	fmt.Scanf("%fn", &monthlySalary)
	if monthlySalary <= 30000 {
		totalPayableTax := monthlySalary * 0.15
		fmt.Printf("Your payable tax for this month is %.2f\n", totalPayableTax)
	}
	if monthlySalary > 30000 {
		totalPayableTax := monthlySalary * 0.20
		fmt.Printf("Your payable tax for this month is %.2f\n", totalPayableTax)
	}
}
