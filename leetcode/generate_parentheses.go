package leetcode

// https://leetcode.com/problems/generate-parentheses/
//
// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
//
// For example, given n = 3, a solution set is:
//
// [
//   "((()))",
//   "(()())",
//   "(())()",
//   "()(())",
//   "()()()"
// ]

func generateParenthesis(n int) []string {
	res := []string{}
	if n == 0 {
		return res
	}
	backtrack(&res, "", 0, 0, n)
	return res
}

func backtrack(res *[]string, s string, open, close, max int) {
	if len(s) == max*2 {
		*res = append(*res, s)
		return
	}
	if open < max {
		backtrack(res, s+"(", open+1, close, max)
	}
	if close < open {
		backtrack(res, s+")", open, close+1, max)
	}
}