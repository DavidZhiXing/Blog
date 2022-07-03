public int Search(int[] nums, int target) {
    int left = 0;
    int right = nums.Length - 1;
    while (left <= right) {
        int mid = left + (right - left) / 2;
        if (nums[mid] == target) {
            return mid;
        }
        if (nums[left] <= nums[mid]) {
            if (nums[left] <= target && target < nums[mid]) {
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        } else {
            if (nums[mid] < target && target <= nums[right]) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
    }
    return -1;
}

// test
// [4,5,6,7,0,1,2]
// 0
// [4,5,6,7,0,1,2]
// 3

[Fact]
public void Test()
{
    var nums = new int[] { 4, 5, 6, 7, 0, 1, 2 };
    var target = 0;
    Assert.Equal(0, Search(nums, target));
}