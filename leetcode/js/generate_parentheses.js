const generateParenthesis = function(n) {
    let res = [];
    backtrack(res, '', 0, 0, n);
    return res;
};

const backtrack = function(res, cur, open, close, max) {
    if (cur.length === max * 2) {
        res.push(cur);
        return;
    }
    if (open < max) {
        backtrack(res, cur + '(', open + 1, close, max);
    }
    if (close < open) {
        backtrack(res, cur + ')', open, close + 1, max);
    }
};

//test
console.log(generateParenthesis(3));
console.log(generateParenthesis(4));
console.log(generateParenthesis(5));