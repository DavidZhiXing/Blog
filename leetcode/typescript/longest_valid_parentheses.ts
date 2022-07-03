let LongestValidParentheses = function() {
    this.longestValidParentheses = function(s) {
        let stack = [];
        let max = 0;
        for (let i = 0; i < s.length; i++) {
        if (s[i] == "(" || s[i] == "[" || s[i] == "{") {
            stack.push(i);
        } else {
            if (stack.length == 0) {
            continue;
            }
            let index = stack.pop();
            if (s[index] == "(" && s[i] == ")" || s[index] == "[" && s[i] == "]" || s[index] == "{" && s[i] == "}") {
            max = Math.max(max, i - index + 1);
            } else {
            stack.push(index);
            }
        }
        }
        return max;
    };
};

// test
let longestValidParentheses = new LongestValidParentheses();
console.log(longestValidParentheses.longestValidParentheses("(()"));
console.log(longestValidParentheses.longestValidParentheses("()(()"));
console.log(longestValidParentheses.longestValidParentheses("()(()"));
console.log(longestValidParentheses.longestValidParentheses("()(()"));
console.log(longestValidParentheses.longestValidParentheses("()(()"));
    
