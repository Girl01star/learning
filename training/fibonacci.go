package training

import (
	"strconv"
)

func FibonacciIterative(n int) int {
	if n < 0 {
		return n
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1

	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func FibonacciRecursive(n int) int {
	if n < 0 {
		return n
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func IsBinaryPalindrome(n int) bool {
	if n < 0 {
		return false
	}
	s := strconv.FormatInt(int64(n), 2)

	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func ValidParentheses(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
			continue
		}
		if ch == ')' || ch == ']' || ch == '}' {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if top != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
			continue
		}
		return false
	}
	return len(stack) == 0
}

func Increment(num string) int {
	runes := []rune(num)
	length := len(runes)
	result := 0

	for i := range length {
		if runes[length-1-i] == '1' {
			result += 1 << i
		}
	}

	return result + 1
}
