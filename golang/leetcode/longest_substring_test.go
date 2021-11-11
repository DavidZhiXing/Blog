package leetcode

import (
	"testing"
)

type testCase struct {
	s      string
	expect int
}

func TestLongestSubstring(t *testing.T) {
	cases := []testCase{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"a", 1},
		{"ab", 2},
		{"abc", 3},
		{"abcd", 4},
		{"abcde", 5},
		{"abcdef", 6},
		{"abcdefg", 7},
		{"abcdefgh", 8},
		{"abcdefghi", 9},
		{"abcdefghij", 10},
		{"abcdefghijk", 11},
		{"abcdefghijkl", 12},
		{"abcdefghijklm", 13},
		{"abcdefghijklmn", 14},
		{"abcdefghijklmno", 15},
		{"abcdefghijklmnop", 16},
		{"abcdefghijklmnopq", 17},
		{"abcdefghijklmnopqr", 18},
		{"abcdefghijklmnopqrs", 19},
		{"abcdefghijklmnopqrst", 20},
		{"abcdefghijklmnopqrstu", 21},
		{"abcdefghijklmnopqrstuv", 22},
		{"abcdefghijklmnopqrstuvw", 23},
		{"abcdefghijklmnopqrstuvwx", 24},
		{"abcdefghijklmnopqrstuvwxy", 25},
		{"abcdefghijklmnopqrstuvwxyz", 26},
		{"abcdefghijklmnopqrstuvwxyza", 27},
		{"abcdefghijklmnopqrstuvwxyzab", 28},
		{"abcdefghijklmnopqrstuvwxyzabc", 29},
		{"abcdefghijklmnopqrstuvwxyzabcd", 30},
		{"abcdefghijklmnopqrstuvwxyzabcde", 31},
	}

	for _, c := range cases {
		if got := lengthOfLongestSubstring(c.s); got != c.expect {
			t.Errorf("lengthOfLongestSubstring(%s) = %d; expect %d", c.s, got, c.expect)
		}
	}
}