package main

import (
	"math"
)

/*

	A palindromic number reads the same both ways. The largest palindrome made
	from the product of two 2-digit numbers is 9009 = 91 × 99.

	Find the largest palindrome made from the product of two 3-digit numbers.

*/

const START int = 999
const END int = 100

func main() {
	palindrome  := -1
	i := START
	ii := START
	for {
		if ii <= END {
			break
		}
		p, ok := isPalindrome(i*ii)
		if ok {
			if palindrome < p {
				palindrome = p
			}
		}
		if i == END {
			ii--
			i = ii
		}
		i--
	}
	println("The answer is: ", palindrome)
}

func isPalindrome(num int) (int, bool) {
	limit := int(math.Log10(float64(num)))
	newNum := 0
	max := int(math.Pow(10, float64(limit)))
	for i := limit; i >= 0; i-- {
		ii := int(math.Pow(10, float64(i)))
		n := (num/ii) % 10
		newNum += n * (max/ii)
	}
	return newNum, newNum == num
}
