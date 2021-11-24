package leetcode

// Given a sorted array, remove the duplicates in place such that each element appear only once and return the new length.
//
// Do not allocate extra space for another array, you must do this in place with constant memory.
//
// For example,
// Given input array nums = [1,1,2],
//
// Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively. It doesn't matter what you leave beyond the new length.

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func main() {
	nums := []int{1, 1, 2}
	n := removeDuplicates(nums)
	fmt.Println(n)
}