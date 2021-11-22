package leetcode

// Given a string s containing just the characters '(', ')', find the length of the longest valid (well-formed) parentheses substring.
//
// For "(()", the longest valid parentheses substring is "()", which has length = 2.
//
// Another example is ")()())", where the longest valid parentheses substring is "()()", which has length = 4.
//
// Example
//
// Given s = ")()())", return 4.
//

func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	stack := make([]int, 0)
	stack = append(stack, -1)
	max := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				max = maxInt(max, i-stack[len(stack)-1])
			}
		}
	}
	return max
}
