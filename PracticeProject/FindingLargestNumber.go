package main

import "fmt"

func main() {
	largest := 0
	for i := 0; i <= 10; i++ {
		fmt.Println("Enter number: ")
		var num int
		fmt.Scanf("%d", &num)
		if num > largest {
			largest = i
		}
	}
	fmt.Println(largest)
}
