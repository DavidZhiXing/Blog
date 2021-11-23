fn searchInRotatedSortedList(nums: Vec<i32>, target: i32) -> i32 {
    let mut l = 0;
    let mut r = nums.len() - 1;
    while l <= r {
        let m = (l + r) / 2;
        if nums[m] == target {
            return m as i32;
        }
        if nums[l] <= nums[m] {
            if nums[l] <= target && target < nums[m] {
                r = m - 1;
            } else {
                l = m + 1;
            }
        } else {
            if nums[m] < target && target <= nums[r] {
                l = m + 1;
            } else {
                r = m - 1;
            }
        }
    }
    -1
}

fn main() {
    let nums = vec![4, 5, 6, 7, 0, 1, 2];
    let target = 0;
    let result = searchInRotatedSortedList(nums, target);
    println!("{}", result);
}