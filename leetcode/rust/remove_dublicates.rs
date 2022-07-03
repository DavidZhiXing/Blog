fn removeDuplicates(nums: &mut Vec<i32>) -> i32 {
    let mut i = 0;
    for j in 1..nums.len() {
        if nums[i] != nums[j] {
            i += 1;
            nums[i] = nums[j];
        }
    }
    (i + 1) as i32
}

// test
fn main() {
    let mut nums = vec![1, 1, 2];
    let res = removeDuplicates(&mut nums);
    println!("{:?}", nums);
    println!("{}", res);
}