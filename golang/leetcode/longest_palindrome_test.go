package leetcode

import (
	"testing"
)

type testCase struct {
	input  string
	output string
}

func TestLongestSubstringPalindrome(t *testing.T) {
	cases := []testCase{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"", ""},
		{"aab", "aa"},
		{"babadada", "adada"},
	}
	for _, c := range cases {
		if longestPalindrome(c.input) != c.output {
			t.Fail()
		}
	}
}