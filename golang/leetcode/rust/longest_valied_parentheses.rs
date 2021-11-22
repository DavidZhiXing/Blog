fn longestValidParentheses(s: String) -> i32 {
    let mut stack: Vec<usize> = Vec::new();
    let mut max_len = 0;
    let mut i = 0;
    while i < s.len() {
        if s.as_bytes()[i] == b'(' {
            stack.push(i);
        } else {
            if stack.len() > 0 {
                stack.pop();
                max_len = max(max_len, i - stack.last().unwrap_or(&0) + 1);
            } else {
                stack.push(i);
            }
        }
        i += 1;
    }
    max_len
}

#[test]
fn test_longestValidParentheses() {
    let s = "()".to_string();
    assert_eq!(longestValidParentheses(s), 2);
    let s = "(()".to_string();
    assert_eq!(longestValidParentheses(s), 2);
    let s = ")()())".to_string();
    assert_eq!(longestValidParentheses(s), 4);

}