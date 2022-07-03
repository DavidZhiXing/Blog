fn InsertPosition(nums: Vec<i32>, target: i32) -> i32 {
    let mut l = 0;
    let mut r = nums.len() - 1;
    let mut m = 0;
    if nums[l] >= target {
        return 0
    }
    while l <= r {
        m = l + (r - l)/2;
        if nums[m] >= target {
            r = m - 1;
        } else {
            if (m == (nums.len() - 1)) || (nums[m + 1] >= target) {
                return m as i32 + 1;
            } 
            l = m + 1;
        }
    }
    return 0
}


#[test]
fn test_InsertPosition() {
    assert_eq!(InsertPosition(vec![1, 3, 5, 6], 5), 2);
    assert_eq!(InsertPosition(vec![1, 3, 5, 6], 2), 1);
    assert_eq!(InsertPosition(vec![1, 3, 5, 6], 7), 4);
    assert_eq!(InsertPosition(vec![1, 3, 5, 6], 0), 0);
    assert_eq!(InsertPosition(vec![1, 3, 5, 6], 6), 3);
    assert_eq!(InsertPosition(vec![1, 3, 5, 6], 4), 2);
    assert_eq!(InsertPosition(vec![1, 3, 5, 6], 1), 0);
    assert_eq!(InsertPosition(vec![1, 3, 5, 6], 8), 4);
    assert_eq!(InsertPosition(vec![1, 3, 5, 6], 9), 4);
    assert_eq!(InsertPosition(vec![1, 3, 5, 6], 10), 4);
    assert_eq!(InsertPosition(vec![1, 3, 5, 6], -1), 0);
    assert_eq!(InsertPosition(vec![1, 3, 5, 6], -2), 0);
    assert_eq!(InsertPosition(vec![1, 3, 5, 6], -3), 0);
    assert_eq!(InsertPosition(vec![1, 3, 5, 6], -4), 0);
    assert_eq!(InsertPosition(vec![1, 3, 5, 6], -5), 0);
    assert_eq!(InsertPosition(vec![1, 3, 5, 6], -6), 0);
    assert_eq!(InsertPosition(vec![1, 3, 5, 6], -7), 0);
    assert_eq!(InsertPosition(vec![1, 3, 5, 6], -8), 0);
    assert_eq!(InsertPosition(vec![1, 3, 5, 6], -9), 0);
}
