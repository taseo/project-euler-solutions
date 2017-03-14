package main

import "fmt"

func main() {

	var n, result int

	fmt.Print("Function finds the sum of all the multiples of 3 or 5 below the provided number\nPlease enter the number: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			result += i
		}
	}

	fmt.Printf("The sum of all the multiples of 3 or 5 below %d is %d\n", n, result)
}
