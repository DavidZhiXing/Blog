package leetcode

import (
	"testing"
)

type TestCase struct {
	input  int
	output int
}

func TestRevertIneger(){
	var testCases = []TestCase{
		{123, 321},
		{-123, -321},
		{120, 21},
		{0, 0},
	}

	for _, testCase := range testCases {
		if revertInteger(testCase.input) != testCase.output {
			t.Errorf("RevertInteger(%d) != %d", testCase.input, testCase.output)
		}
	}
}
