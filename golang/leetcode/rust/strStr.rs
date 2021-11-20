fn strStr(haystack: String, needle: String) -> i32 {
    let haystack_bytes = haystack.as_bytes();
    let needle_bytes = needle.as_bytes();
    let mut i = 0;
    let mut j = 0;
    while i < haystack_bytes.len() && j < needle_bytes.len() {
        if haystack_bytes[i] != needle_bytes[j] {
            i += 1;
        } else {
            i += 1;
            j += 1;
        }
    }
    if j == needle_bytes.len() {
        return i - j;
    } else {
        return -1;
    }
}

// test
fn main() {
    assert_eq!(strStr("hello".to_string(), "ll".to_string()), 2);
    assert_eq!(strStr("aaaaa".to_string(), "bba".to_string()), -1);
    assert_eq!(strStr("aaaaa".to_string(), "aaaaa".to_string()), 0);
    assert_eq!(strStr("".to_string(), "".to_string()), 0);
    assert_eq!(strStr("".to_string(), "a".to_string()), -1);
    assert_eq!(strStr("a".to_string(), "".to_string()), 0);
    assert_eq!(strStr("a".to_string(), "a".to_string()), 0);
    assert_eq!(strStr("a".to_string(), "b".to_string()), -1);
    assert_eq!(strStr("b".to_string(), "a".to_string()), -1);
    assert_eq!(strStr("b".to_string(), "b".to_string()), 0);
    assert_eq!(strStr("b".to_string(), "c".to_string()), -1);
    assert_eq!(strStr("c".to_string(), "b".to_string()), -1);
    assert_eq!(strStr("c".to_string(), "c".to_string()), 0);
    assert_eq!(strStr("c".to_string(), "d".to_string()), -1);
    assert_eq!(strStr("d".to_string(), "c".to_string()), -1);
    assert_eq!(strStr("d".to_string(), "d".to_string()), 0);
    assert_eq!(strStr("d".to_string(), "e".to_string()), -1);
    assert_eq!(strStr("e".to_string(), "d".to_string()), -1);
    assert_eq!(strStr("e".to_string(), "e".to_string()), 0);
}