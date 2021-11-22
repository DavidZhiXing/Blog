let next_Permutation = function(nums) {
    let count = 0;
    let start = 0;
    let end = nums.length - 1;
    while (start < end) {
        if (nums[start] < nums[end]) {
            count = end - start;
            break;
        }
        start++;
    }
    if (count === 0) {
        return nums.reverse();
    }
    let index = start;
    while (index < end) {
        if (nums[index] > nums[start]) {
            break;
        }
        index++;
    }
    let temp = nums[start];
    nums[start] = nums[index];
    nums[index] = temp;
    let left = start + 1;
    let right = end;
    while (left < right) {
        temp = nums[left];
        nums[left] = nums[right];
        nums[right] = temp;
        left++;
        right--;
    }
    return nums;
};

