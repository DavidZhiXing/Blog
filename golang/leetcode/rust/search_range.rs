fn searchRange( nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut l = 0;
    let mut r = nums.len() - 1;
    let mut m = 0;
    while l <= r {
        m = (l + r) / 2;
        if nums[m] == target {
            return vec![m as i32, m as i32];
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
    vec![-1, -1]
}