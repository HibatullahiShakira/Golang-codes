package main

import "fmt"

func main() {
	var grade string
	var countA int = 0
	var countB int = 0
	var countC int = 0
	var countD int = 0
	for i := 0; i < 5; i++ {
		fmt.Println("Enter student grade")
		fmt.Scan(&grade)

		switch grade {
		case "A":
			countA++
		case "B":
			countB++
		case "C":
			countC++
		case "D":
			countD++
		}
	}
	fmt.Println("The total number of student with grade A is ", countA,
		"\n The total number of student with grade B is ", countB,
		"\n The total number of student with grade C is ", countC,
		"\n The total number of student with grade D is ", countD)
}
