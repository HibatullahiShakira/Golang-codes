package main

import "fmt"

func main() {
	fmt.Println("Enter the number of miles driven")
	var milesDriven float64
	fmt.Scanf("%f", &milesDriven)
	fmt.Println("Enter the number of gallon used")
	var gallonUsed float64
	fmt.Scanf("%f", &gallonUsed)
	totalMilesPerGallon := milesDriven / gallonUsed
	fmt.Printf("The total miles driven is per gallon is %.2fn\n", totalMilesPerGallon)
}
