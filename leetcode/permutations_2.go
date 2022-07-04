// given a collection of numbersw that might contain duplicates, return all possible unique permutations.
// example:
// Input: [1,1,2]
// Output:
// [
//   [1,1,2],
//   [1,2,1],
//   [2,1,1]
// ]

func permutations_2(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int()
	}
	used, p, res := make([]bool, len(nums)), []int{}, [][]int{}
	generate_2(nums, 0, p, &res, &used)
	return res
}

func generate_2(nums []int, index int, p []int, res *[][]int, used *[]bool) {
	if index == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, p)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			// distinct number
			if i > 0 && nums[i] == nums[i-1] && !(*used)[i-1] {
				continue
			}
			(*used)[i] = true
			p = append(p, nums[i])
			generate_2(nums, index+1, p, res, used)
			p = p[:len(p)-1]
			(*used)[i] = false
		}
	}
}