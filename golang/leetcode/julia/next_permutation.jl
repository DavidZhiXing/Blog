function next_permutation(nums::Arr)
    if nums.length <= 1
        return nums
    end
    i = nums.length - 2
    while i >= 0
        if nums[i] < nums[i+1]
            j = nums.length - 1
            while nums[i] >= nums[j]
                j -= 1
            end
            nums[i], nums[j] = nums[j], nums[i]
            nums[i+1..-1].reverse!
            return nums
        end
        i -= 1
    end
    nums.reverse!
    return nums
end
#### test
# p next_permutation([1,2,3])
# p next_permutation([3,2,1])
# p next_permutation([1,1,5])
# p next_permutation([1,5,1])
# p next_permutation([1])
# p next_permutation([])
