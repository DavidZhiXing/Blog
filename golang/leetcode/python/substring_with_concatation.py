class LeetCode:
    def substringWithConcatationOfAllWords(self, s, words):
        """
        :type s: str
        :type words: List[str]
        :rtype: List[int]
        """
        if not s or not words:
            return []
        if len(words) == 1:
            return [s.find(words[0])] if words[0] in s else []
        words.sort(key=lambda x: len(x))
        res = []
        for i in range(len(s)):
            for j in range(len(words)):
                if i + len(words[j]) > len(s):
                    break
                if s[i:i + len(words[j])] == words[j]:
                    if j == len(words) - 1:
                        res.append(i)
                    if j < len(words) - 1:
                        for k in range(i + len(words[j]), len(s)):
                            if s[k:k + len(words[j + 1])] == words[j + 1]:
                                res.append(i)
                                break
        return res

# Test
if __name__ == '__main__':
    leetcode = LeetCode()
    print(leetcode.substringWithConcatationOfAllWords("barfoothefoobarman", ["foo", "bar"]))
    print(leetcode.substringWithConcatationOfAllWords("wordgoodgoodgoodbestword", ["word", "good", "best", "word"]))
    print(leetcode.substringWithConcatationOfAllWords("barfoofoobarthefoobarman", ["bar", "foo", "the"]))
    print(leetcode.substringWithConcatationOfAllWords("wordgoodgoodgoodbestword", ["word", "good", "best", "good"]))
    print(leetcode.substringWithConcatationOfAllWords("barfoofoobarthefoobarman", ["bar", "foo", "the", "foo"]))
    print(leetcode.substringWithConcatationOfAllWords("barfoofoobarthefoobarman", ["bar", "foo", "the", "foobar"]))
    print(leetcode.substringWithConcatationOfAllWords("barfoofoobarthefoobarman", ["bar", "foo", "the", "foobar", "