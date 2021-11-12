package leetcode

import (
	"testing"
)

type testCase struct {
	input  string
	output string
}

func TestZigZag(t *testing.T) {
	testCases := []testCase{
		{"PAYPALISHIRING", "PAHNAPLSIIGYIR"},
		{"ABCDE", "ABCDE"},
		{"ABCDEFG", "AGCEDBF"},
		{"A", "A"},
		{"", ""},
		{"Z", "Z"},
		{"ZABCDEFG", "ZAGCEDBF"},
	}

	for _, testCase := range testCases {
		if zigZag(testCase.input) != testCase.output {
			t.Errorf("zigZag(%s) == %s, expected %s", testCase.input, zigZag(testCase.input), testCase.output)
		}
	}
}

