package main

import "fmt"

func main() {
	fmt.Println("FibonacciIterative(10):", FibonacciIterative(10))

	fmt.Println("FibonacciRecursive(10):", FibonacciRecursive(10))

	fmt.Println("IsPrime(2):", IsPrime(2))
	fmt.Println("IsPrime(15):", IsPrime(15))
	fmt.Println("IsPrime(29):", IsPrime(29))

	fmt.Println("IsBinaryPalindrome(7):", IsBinaryPalindrome(7))
	fmt.Println("IsBinaryPalindrome(6):", IsBinaryPalindrome(6))

	fmt.Println(`ValidParentheses("[]{}()"):`, ValidParentheses("[]{}()"))
	fmt.Println(`ValidParentheses("[{]}"):`, ValidParentheses("[{]}"))

	fmt.Println(`Increment("101") ->`, Increment("101"))

}
