// Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
// example:
// Input: "babad"
// Output: "bab"
// Note: "aba" is also a valid answer.

// example:
// Input: "cbbd"
// Output: "bb"

package leetcode

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		len := max(len1, len2)
		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}
	return s[start : end+1]
}

func longestPalindromeDP(s string) string {
	if len(s) < 2 {
		return s
	}
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if s[i] == s[j] && (i-j < 2 || dp[j+1][i-1]) {
				dp[j][i] = true
				if i-j > end-start {
					start = j
					end = i
				}
			}
		}
	}
	return s[start : end+1]
}	



func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestPalindromeManacher(s string) string {
	if len(s) < 2 {
		return s
	}
	t := preProcess(s)
	p := make([]int, len(t))
	c := 0
	r := 0
	for i := 1; i < len(t)-1; i++ {
		if i < r {
			p[i] = min(r-i, p[2*c-i])
		}
		for i+p[i]+1 < len(t) && t[i+p[i]+1] == t[i-p[i]-1] {
			p[i]++
		}
		if i+p[i] > r {
			c = i
			r = i + p[i]
		}
	}
	maxLen := 0
	centerIndex := 0
	for i := 0; i < len(p); i++ {
		if p[i] > maxLen {
			maxLen = p[i]
			centerIndex = i
		}
	}
	start := (centerIndex - maxLen) / 2
	return s[start : start+maxLen]
}

func preProcess(s string) string {
	t := "#"
	for i := 0; i < len(s); i++ {
		t += string(s[i]) + "#"
	}
	return t
}