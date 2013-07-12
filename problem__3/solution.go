package main

import (
	"math/big"
)

/*

	Problem:

	The prime factors of 13195 are 5, 7, 13 and 29.

	What is the largest prime factor of the number 600851475143 ?

*/

type Node struct {
	next  *Node
	value *big.Int
}

func NewNode(value *big.Int) *Node {
	return &Node{
		value: value,
		next:  nil,
	}
}

func (n *Node) Last() *Node {
	if n == nil {
		return nil
	}
	current := n
	for current.next != nil {
		current = current.next
	}
	return current
}

const INITIAL_CEIL int64 = 100

var primeFactors *Node
var problem *big.Int
var currentPrimeCeil *big.Int = big.NewInt(INITIAL_CEIL)
var primeCeilIncrease *big.Int = big.NewInt(10)
var zero *big.Int = big.NewInt(0)
var two *big.Int = big.NewInt(2)

const PROBLEM string = "600851475143"

func main() {
	problem = big.NewInt(0)
	_, ok := problem.SetString(PROBLEM, 10)
	if !ok {
		println("Not ok setting problem with big number operation.")
		return
	}
	start := problem
	findPrimeFactors(start, currentPrimeCeil)
	println("The answer is: ", primeFactors.Last().value.String())
}

func findPrimeFactors(num, ceil *big.Int) {
	// See if we have answer
	current := primeFactors
	value := big.NewInt(1)
	for current != nil {
		value.Mul(value, current.value)
		current = current.next
	}
	if value.Cmp(problem) == 0 {
		return
	}
	// Try 2 first.
	quotient, ok := isFactor(num, two)
	if ok {
		addToPrimeFactors(two)
		findPrimeFactors(quotient, ceil)
		return
	}
	head := NewNode(big.NewInt(3))
	current = head
	for index := big.NewInt(5); index.Cmp(ceil) != 1; index.Add(index, two) {
		value := big.NewInt(0)
		value.Set(index)
		node := NewNode(value)
		current.next = node
		current = node
	}
	prime := big.NewInt(2)
	nextQuotient := filterPrimes(head, num, prime)
	if nextQuotient.Cmp(num) == 0 {
		// Try again with a higher prime factor search.
		ceil.Mul(ceil, primeCeilIncrease)
	} else {
		// Found a new quotient, reset prime factor ceiling.
		ceil.SetInt64(int64(INITIAL_CEIL))
	}
	findPrimeFactors(nextQuotient, ceil)
}

func isFactor(num, div *big.Int) (next *big.Int, dividedEvenly bool) {
	rnum := big.NewRat(0, 10)
	rnum.SetInt(num)
	rdiv := big.NewRat(0, 10)
	rdiv.SetInt(div)
	rquotient := big.NewRat(0, 10)
	rquotient.Quo(rnum, rdiv)
	if !rquotient.IsInt() {
		return num, false
	}
	quotient := big.NewInt(0)
	quotient.Div(num, div)
	value := big.NewInt(0)
	value.Mod(num, quotient)
	if value.Cmp(zero) == 0 {
		return quotient, true
	}
	return num, false
}

func addToPrimeFactors(val *big.Int) {
	value := big.NewInt(0)
	value.Set(val)
	if primeFactors == nil {
		primeFactors = NewNode(val)
	} else {
		current := primeFactors.Last()
		current.next = NewNode(val)
	}
}

func filterPrimes(head *Node, num, prime *big.Int) *big.Int {
	if head == nil {
		return num
	}
	prime.Set(head.value)
	quotient, ok := isFactor(num, prime)
	if ok {
		addToPrimeFactors(prime)
		return quotient
	}
	current := head
	for current.next != nil {
		test := big.NewInt(0)
		test.Set(current.next.value)
		value := big.NewInt(0)
		value = value.Mod(test, prime)
		if value.Cmp(zero) == 0 {
			current.next = current.next.next
		} else {
			current = current.next
		}
	}
	return filterPrimes(head.next, num, prime)
}
