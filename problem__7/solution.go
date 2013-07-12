package main

/*

	Problem:

	By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see
	that the 6th prime is 13.

	What is the 10001st prime number?

*/

const numPrimes uint64 = 10001
const limit uint64 = 200000

var current []uint64
var lastPrime uint64 = 2
var foundPrimes uint64 = 1

func main() {
	current = make([]uint64, limit+1)
	count := 0
	for i := uint64(3); i < limit+1; i += 2 {
		current[count] = i
		count++
	}
	filter(current[:count])
	println("The answer is: ", lastPrime)
}

func filter(numbers []uint64) {
	var num uint64
	if foundPrimes == numPrimes {
		return
	}
	lastPrime = numbers[0]
	count := 0
	for i := 0; i < len(numbers); i++ {
		num = numbers[i]
		if num%lastPrime > 0 {
			current[count] = num
			count++
		}
	}
	foundPrimes++
	filter(current[:count])
}
