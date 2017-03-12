package main

import (
	"fmt"
	"math/big"
)

func main() {

	var n int

	finder := func(n int64) {

		// Initialize first two Fibonacci sequence numbers and  index position

		a, b, index := big.NewInt(1), big.NewInt(1), big.NewInt(1)

		// Initialize limit as 10^n-1 (e.g., the smallest integer with 100 digits would be 10^99)

		var limit big.Int
		limit.Exp(big.NewInt(10), big.NewInt(n-1), nil)

		// Loop while a is smaller than limit (e.g., smaller than 1e100 for 100 digit number)

		for a.Cmp(&limit) < 0 {

			// Compute the next Fibonacci number, swamp a an b so that b is next number in the sequence and icrement index

			a.Add(a, b)

			a, b = b, a

			index.Add(index, big.NewInt(1))
		}

		fmt.Printf("The first %d digit number in Fibonacci sequence with index F%d is %d\n", n, index, a)
	}

	fmt.Print("Functions finds first number in Fibonacci sequence and its position (index) based on digit count in it\nHow many digits the number should have: ")
	fmt.Scan(&n)
	finder(int64(n))
}
