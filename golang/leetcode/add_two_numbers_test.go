package leetcode

import (
	"testing"
	"fmt"
	"structures"
)

type addTwoNumbersTest struct {
	num1 *structures.ListNode
	num2 *structures.ListNode
	expected *structures.ListNode
}

type addTwoNumbersTestCase struct {
	testCase addTwoNumbersTest
	description string
}

type addTwoNumbersTestCases []addTwoNumbersTestCase


func addTwoNumbersTestCaseProvider(t *testing.T) addTwoNumbersTestCases {
	return []addTwoNumbersTestCase{
		{
			testCase: addTwoNumbersTest{
				num1: structures.ListNodeFromSlice([]int{2, 4, 3}),
				num2: structures.ListNodeFromSlice([]int{5, 6, 4}),
				expected: structures.ListNodeFromSlice([]int{7, 0, 8}),
			},
			description: "Two numbers of different length",
		},
		{
			testCase: addTwoNumbersTest{
				num1: structures.ListNodeFromSlice([]int{5}),
				num2: structures.ListNodeFromSlice([]int{5}),
				expected: structures.ListNodeFromSlice([]int{0, 1}),
			},
			description: "Two numbers of same length",
		},
		{
			testCase: addTwoNumbersTest{
				num1: structures.ListNodeFromSlice([]int{5}),
				num2: structures.ListNodeFromSlice([]int{5, 6}),
				expected: structures.ListNodeFromSlice([]int{0, 1, 1}),
			},
			description: "Two numbers of same length with carry",
		},
	}
}

func TestAddTwoNumbers(t *testing.T) {
	testCases := addTwoNumbersTestCaseProvider(t)

	for _, testCase := range testCases {
		result := addTwoNumbers(testCase.testCase.num1, testCase.testCase.num2)

		if !structures.ListNodesEqual(result, testCase.testCase.expected) {
			t.Errorf("%s: expected %v, got %v", testCase.description, testCase.testCase.expected, result)
		}
	}
}