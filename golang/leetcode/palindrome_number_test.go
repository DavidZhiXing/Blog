package leetcode 

import (
	"testing"
)

type TestCase struct {
	input  int
	output bool
}

func TestIsPalindrome(t *testing.T) {
	var testCases = []TestCase{
		{121, true},
		{-121, false},
		{10, false},
	}

	for _, testCase := range testCases {
		if isPalindrome(testCase.input) != testCase.output {
			t.Errorf("isPalindrome(%d) = %t, expected %t", testCase.input, isPalindrome(testCase.input), testCase.output)
		}
	}
}