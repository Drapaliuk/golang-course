package main

import (
	"fmt"
	"lesson-2/binaryPalindrome"
	"lesson-2/fibonacci"
	"lesson-2/increment"
	"lesson-2/primeNumbers"
	"lesson-2/validParentheses"
)

func main() {
	// fibonacci iterative
	fmt.Println("Fibonacci iterative")
	fibIter1 := fibonacci.FibonacciIterative(3)
	fibIter2 := fibonacci.FibonacciIterative(10)
	fibIter3 := fibonacci.FibonacciIterative(12)
	fmt.Println("FibIter_1: ", fibIter1)
	fmt.Println("FibIter_2: ", fibIter2)
	fmt.Println("FibIter_3: ", fibIter3)
	fmt.Println("\n")

	// fibonacci recursive
	fmt.Println("Fibonacci recursive")
	fibRecurs1 := fibonacci.FibonacciRecursive(3)
	fibRecurs2 := fibonacci.FibonacciRecursive(10)
	fibRecurs3 := fibonacci.FibonacciRecursive(12)
	fmt.Println("FibRecurs_1: ", fibRecurs1)
	fmt.Println("FibRecurs_2: ", fibRecurs2)
	fmt.Println("FibRecurs_3: ", fibRecurs3)
	fmt.Println("\n")

	// bindary palindrome
	fmt.Println("Bindary palindrome")
	binPalindrome1 := binaryPalindrome.IsBinaryPalindrome(7) // 7 -> 111 -> true
	binPalindrome2 := binaryPalindrome.IsBinaryPalindrome(5) // 5 -> 101 -> true
	binPalindrome3 := binaryPalindrome.IsBinaryPalindrome(6) // 6 -> 110 -> false
	fmt.Println("BinPalindrome_1: ", binPalindrome1)
	fmt.Println("BinPalindrome_2: ", binPalindrome2)
	fmt.Println("BinPalindrome_3: ", binPalindrome3)
	fmt.Println("\n")

	// valid parentheses
	fmt.Println("Valid parentheses")
	validParentheses_1 := validParentheses.ValidParentheses("[{}]")             // true
	validParentheses_2 := validParentheses.ValidParentheses("[{{{{[[()]]}}}}]") // true
	validParentheses_3 := validParentheses.ValidParentheses("[{]}")             // false
	validParentheses_4 := validParentheses.ValidParentheses("[{(a[]))}}")       //false
	fmt.Println("ValidParentheses_1: ", validParentheses_1)
	fmt.Println("ValidParentheses_2: ", validParentheses_2)
	fmt.Println("ValidParentheses_3: ", validParentheses_3)
	fmt.Println("ValidParentheses_4: ", validParentheses_4)
	fmt.Println("\n")

	// increment
	fmt.Println("Increment")
	increment_1 := increment.Increment("0")         // 0 -> 0
	increment_2 := increment.Increment("101")       // 5 -> 6
	increment_3 := increment.Increment("1100011")   // 99 -> 100
	increment_4 := increment.Increment("101011001") // 345 -> 346
	fmt.Println("Increment_1: ", increment_1)
	fmt.Println("Increment_2: ", increment_2)
	fmt.Println("Increment_3: ", increment_3)
	fmt.Println("Increment_4: ", increment_4)
	fmt.Println("\n")

	// is prime not implemented
	fmt.Println("Increment")
	primeNumbers.IsPrime(2)

}
