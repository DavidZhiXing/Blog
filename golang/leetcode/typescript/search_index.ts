const searchIndexInSortedArray = (nums: number[], target: number): number => {
    let start = 0;
    let end = nums.length - 1;
    while (start <= end) {
        const mid = Math.floor((start + end) / 2);
        if (nums[mid] < target) {
            start = mid + 1;
        } else if (nums[mid] > target) {
            end = mid - 1;
        } else {
            if (mid === 0 || nums[mid - 1] !== target) {
                return mid;
            }
            end = mid - 1;
        }
    }
    return -1;
};

// Test
console.log(searchIndexInSortedArray([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 10));
console.log(searchIndexInSortedArray([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 9));
console.log(searchIndexInSortedArray([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 8));
console.log(searchIndexInSortedArray([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 7));
console.log(searchIndexInSortedArray([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 6));