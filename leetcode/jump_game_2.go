// given an array of non-negative integers nums, you are initially positioned at the first index of the array.
// Each element in the array reperesents your maximum jump length at taht postition.
// your goal is to reach the last index in the minimum nember of jumps.
// you can assume that you can always reach the last index.

// example:
// Input: [2,3,1,1,4]
// Output: 2
// Explanation: The minimum number of jumps to reach the last index is 2.

// Example 2:
// Input: [2, 3, 0, 1, 4]
// Output: 2

// constraints:
// 1 <= nums.length <= 1000
// 0 <= nums[i] <= 10^5

package leetcode

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	needChoose, canReach, step := 0, 0, 0
	for i, x := range nums {
		if i + x > canReach {
			canReach = i + x
			if canReach >= len(nums)-1 {
				return step + 1
			}
		}
		if i == needChoose {
			needChoose = canReach
			step++
		}
	}
	return step
}