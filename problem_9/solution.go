package main

/*

	Problem:

	A Pythagorean triplet is a set of three natural numbers, a < b < c, for
	which,

	a^2 + b^2 = c^2
	For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

	There exists exactly one Pythagorean triplet for which a + b + c = 1000.
	Find the product abc.

*/

const LIMIT uint64 = 1000

func main() {
	answer := uint64(0)
	found := false
	for b := uint64(2); b < LIMIT; b++ {
		if found {
			break
		}
		for a := uint64(1); a < b; a++ {
			c := 1000 - (a + b)
			if (a*a)+(b*b) == (c * c) {
				answer = a * b * c
				found = true
				break
			}
		}
	}
	println("The answer is: ", answer)
}
