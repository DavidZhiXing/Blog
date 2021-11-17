class LeetCode_GenerateParentheses(object):
    def generateParenthesis(self, n):
        """
        :type n: int
        :rtype: List[str]
        """
        if n == 0:
            return []
        res = []
        self.dfs(n, n, '', res)
        return res

    def dfs(self, left, right, path, res):
        if left == 0 and right == 0:
            res.append(path)
        if left > 0:
            self.dfs(left - 1, right, path + '(', res)
        if right > left:
            self.dfs(left, right - 1, path + ')', res)

def main():
    leetcode = LeetCode_GenerateParentheses()
    print(leetcode.generateParenthesis(3))
    