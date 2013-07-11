package main

/*

	Problem:

	2520 is the smallest number that can be divided by each of the numbers from 1
	to 10 without any remainder.

	What is the smallest positive number that is evenly divisible by all of the
	numbers from 1 to 20?

*/

const primeProduct uint64 = 2 * 3 * 5 * 7 * 11 * 13 * 17 * 19

var nonPrimes = []uint64{4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20}

func main() {
	answer := uint64(0)
	for i := primeProduct; true; i += primeProduct {
		passed := true
		for _, nonPrime := range nonPrimes {
			if i%nonPrime != 0 {
				passed = false
				break
			}
		}
		if passed {
			answer = i
			break
		}
	}
	println("The answer is: ", answer)
}
