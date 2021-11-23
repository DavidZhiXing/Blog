const searchRange = (nums: number[], target: number): number[] => {
    const start = findStart(nums, target);
    const end = findEnd(nums, target);
    return [start, end];
    };

const findStart = (nums: number[], target: number): number => {
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

const findEnd = (nums: number[], target: number): number => {
    let start = 0;
    let end = nums.length - 1;
    while (start <= end) {
        const mid = Math.floor((start + end) / 2);
        if (nums[mid] < target) {
            start = mid + 1;
        } else if (nums[mid] > target) {
            end = mid - 1;
        } else {
            if (mid === nums.length - 1 || nums[mid + 1] !== target) {
                return mid;
            }
            start = mid + 1;
        }
    }
    return -1;
};

// test
const nums = [5, 7, 7, 8, 8, 10];
const target = 8;
console.log(searchRange(nums, target));
