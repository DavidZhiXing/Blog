package leetcode

import (
	"testing"
	"fmt"
)

type question1 func() {
	para1
	ans1
}

type para1 struct {
	nums []int
	target int
}

type ans1 struct {
	one []int
}

func Test_Problem1(t *testing.T) {
	qs := []question1 {
		func() {
			nums := []int{2, 7, 11, 15}
			target := 9
			ans := []int{0, 1}
			fmt.Println(twoSum(nums, target))
		},
		func() {
			nums := []int{3, 2, 4}
			target := 6
			ans := []int{1, 2}
			fmt.Println(twoSum(nums, target))
		},
		func() {
			nums := []int{3, 3}
			target := 6
			ans := []int{0, 1}
			fmt.Println(twoSum(nums, target))
		},
	}
	for _, q := range qs {
		q()
	}
}