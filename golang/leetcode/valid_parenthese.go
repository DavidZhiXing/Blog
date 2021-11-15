package leetcode

// https://leetcode.com/problems/valid-parentheses/
//
// Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// An input string is valid if:
//
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Note that an empty string is also considered valid.
//
// Example 1:
//
// Input: "()"
// Output: true
//
// Example 2:
//
// Input: "()[]{}"
// Output: true
//
// Example 3:
//
// Input: "(]"
// Output: false
//
// Example 4:
//
// Input: "([)]"
// Output: false
//
// Example 5:
//
// Input: "{[]}"
// Output: true
//
//
// Constraints:
//
//     1 <= text.length <= 104
//     text consists of parentheses only '()[]{}'.

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]rune, 0)
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if (last == '(' && c == ')') || (last == '{' && c == '}') || (last == '[' && c == ']') {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

