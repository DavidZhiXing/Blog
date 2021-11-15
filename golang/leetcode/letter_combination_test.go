package leetcode

import (
	"testing"
)

type testCase struct {
	digits string
	output []string
}

func TestLetterCombinations(t *testing.T) {
	cases := []testCase{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"", []string{}},
	}

	for _, c := range cases {
		if got := letterCombinations(c.digits); !equal(got, c.output) {
			t.Errorf("letterCombinations(%v) = %v, want %v", c.digits, got, c.output)
		}
	}
}