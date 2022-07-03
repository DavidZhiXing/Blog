fn nextPermutation(nums: &mut Vec<i32>) {
    let mut i = nums.len() - 1;
    while i > 0 {
        if nums[i - 1] < nums[i] {
            break;
        }
        i -= 1;
    }
    if i == 0 {
        nums.reverse();
        return;
    }
    let mut j = nums.len() - 1;
    while j >= 0 {
        if nums[i - 1] < nums[j] {
            break;
        }
        j -= 1;
    }
    nums.swap(i - 1, j);
    nums[i..].reverse();
}


#[test]
fn test_next_permutation() {
    let mut nums = vec![1, 2, 3];
    nextPermutation(&mut nums);
    assert_eq!(nums, vec![1, 3, 2]);
    let mut nums = vec![3, 2, 1];
    nextPermutation(&mut nums);
    assert_eq!(nums, vec![1, 2, 3]);
    let mut nums = vec![1, 1, 5];
    nextPermutation(&mut nums);
    assert_eq!(nums, vec![1, 5, 1]);
    let mut nums = vec![1, 3, 2];
    nextPermutation(&mut nums);
    assert_eq!(nums, vec![2, 1, 3]);
    let mut nums = vec![2, 3, 1];
    nextPermutation(&mut nums);
    assert_eq!(nums, vec![3, 1, 2]);
    let mut nums = vec![1, 1, 5];
    nextPermutation(&mut nums);
    assert_eq!(nums, vec![1, 5, 1]);
    let mut nums = vec![1, 2, 3, 4];
    nextPermutation(&mut nums);
    assert_eq!(nums, vec![1, 2, 4, 3]);
}