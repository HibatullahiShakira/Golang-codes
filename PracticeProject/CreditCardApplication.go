package main

import "fmt"

func main() {
	const creditLimit float64 = 50000
	fmt.Println("Enter your account number")
	var accountNumber string
	fmt.Scan(&accountNumber)
	fmt.Println("Enter account balance at the beginning of the month")
	var beginningBalance float64
	fmt.Scan(&beginningBalance)
	fmt.Println("Enter total item charged by customer this month")
	var totalItemCharged float64
	fmt.Scan(&totalItemCharged)
	fmt.Println("Enter total credit charged applied at the end of the month")
	var totalCredit float64
	fmt.Scan(&totalCredit)
	newBalance := beginningBalance + totalItemCharged - totalCredit
	fmt.Println("Your total balance is", newBalance)
	if newBalance > creditLimit {
		fmt.Println("Credit limit reached")
	} else {
		fmt.Println("Lucky you, you haven't reached your credit limit yet you can still spend more")
	}
}
