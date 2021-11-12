package leetcode

import (
	"testing"
)

type testCase struct {
	input  string
	output int
}

func TestStringToInt(t *testing.T) {
	testCases := []testCase{
		{"42", 42},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
	}

	for _, testCase := range testCases {
		result := myAtoi(testCase.input)
		if result != testCase.output {
			t.Errorf("Expected %d, got %d", testCase.output, result)
		}
	}
}