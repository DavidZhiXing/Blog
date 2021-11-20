package leetcode

// given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.
// return the quotient after dividing dividend by divisor.
//
// Note:
//
// The integer division should truncate toward zero.
//
// Example 1:
//
// Input: dividend = 10, divisor = 3
// Output: 3
// Example 2:
//
// Input: dividend = 7, divisor = -3
// Output: -2
// Note:
//
// Both dividend and divisor will be 32-bit signed integers.
// The divisor will never be 0.
// Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 231 − 1 when the division result overflows.

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return 0
	}
	if dividend == 0 {
		return 0
	}
	if dividend == -2147483648 && divisor == -1 {
		return 2147483647
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend == 2147483648 {
			return -2147483648
		}
		return -dividend
	}

	var res int
	var sign int
	if dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0 {
		sign = 1
	} else {
		sign = -1
	}

	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	for dividend >= divisor {
		res++
		dividend -= divisor
	}
	if sign == 1 {
		return res
	}
	return -res
}