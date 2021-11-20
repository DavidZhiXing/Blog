class LeetCodeSolution:
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if len(nums) == 0:
            return 0
        i = 0
        for j in range(1, len(nums)):
            if nums[i] != nums[j]:
                i += 1
                nums[i] = nums[j]
        return i + 1

# Test Code
if __name__ == '__main__':
    test = LeetCodeSolution()
    print(test.removeDuplicates([1, 1, 2]))
    print(test.removeDuplicates([0, 0, 1, 1, 1, 2, 2, 3, 3, 4]))
    print(test.removeDuplicates([1, 1, 1, 2, 2, 3]))