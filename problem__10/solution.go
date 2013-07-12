package main

/*

	Problem:

	The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

	Find the sum of all the primes below two million.

*/

var current []int

const LIMIT int = 2000000

func main() {
	primes := []int{2}
	if LIMIT <= 1 {
		return
	}
	current = make([]int, LIMIT+1)
	count := 0
	for i := 3; i < LIMIT+1; i += 2 {
		current[count] = i
		count++
	}
	primes = filter(current[:count], primes)
	sum := 0
	for _, prime := range primes {
		sum += prime
	}
	println("The answer is: ", sum)
}

func filter(numbers, primes []int) []int {
	var num int
	if len(numbers) < 1 {
		return primes
	}
	prime := numbers[0]
	count := 0
	for i := 0; i < len(numbers); i++ {
		num = numbers[i]
		if num%prime > 0 {
			current[count] = num
			count++
		}
	}
	primes = append(primes, prime)
	return filter(current[:count], primes)
}
