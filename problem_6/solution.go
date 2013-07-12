package main

import (
	"math"
)

/*

	Problem:

	The sum of the squares of the first ten natural numbers is,

	1^2 + 2^2 + ... + 10^2 = 385

	The square of the sum of the first ten natural numbers is,

	(1 + 2 + ... + 10)^2 = 55^2 = 3025

	Hence the difference between the sum of the squares of the first ten natural
	numbers and the square of the sum is 3025 − 385 = 2640.

	Find the difference between the sum of the squares of the first one hundred
	natural numbers and the square of the sum.

*/

const LIMIT uint64 = 100

func main() {
	sumOfSquares := uint64(0)
	sum := uint64(0)
	for i := uint64(1); i <= LIMIT; i++ {
		sumOfSquares += uint64(math.Pow(float64(i), 2))
		sum += i
	}
	squareOfSums := uint64(math.Pow(float64(sum), 2))
	println("The answer is: ", (squareOfSums - sumOfSquares))
}
