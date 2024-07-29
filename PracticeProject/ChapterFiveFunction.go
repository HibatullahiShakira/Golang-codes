package main

import "fmt"

func findMinimumMaximumNumber(arr ...int) int {
	maximum := 0
	minimum := 0
	var arrKeeper []int
	for _, elem := range arr {
		arrKeeper = append(arrKeeper, elem)
		if elem < minimum {
			minimum = elem
		}
		if elem > maximum {
			maximum = elem
		}
	}
	sum := maximum + minimum
	return sum
}

func sumIntegerDivisibleByThree(variable []int) int {
	for i := 0; i < len(variable); i++ {
		if variable[i]%3 == 0 {
			return variable[i]
		}
	}
	return 0
}

func rightAngleTriangle() {
	for i := 0; i < 5; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	//for i := 5; i > 0; i-- {
	//	for j := 0; j <= i; j++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}
}
