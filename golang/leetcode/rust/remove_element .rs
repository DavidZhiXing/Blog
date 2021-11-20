fn removeElement(nums: &mut Vec<i32>, val: i32) -> i32 {
    let mut i = 0;
    for j in 0..nums.len() {
        if nums[j] != val {
            i += 1;
            nums[i] = nums[j];
        }
    }
    i + 1
}

// test
fn main() {
    let mut nums = vec![3,2,2,3];
    let val = 3;
    let result = removeElement(&mut nums, val);
    println!("{:?}", nums);
    println!("{}", result);
}