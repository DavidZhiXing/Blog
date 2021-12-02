def nextPermutation(nums)
  i = nums.length - 2
  while i >= 0 && nums[i] >= nums[i + 1]
    i -= 1
  end
  if i >= 0
    j = nums.length - 1
    while j >= 0 && nums[j] <= nums[i]
      j -= 1
    end
    nums[i], nums[j] = nums[j], nums[i]
  end
  nums[i + 1..-1].reverse!
end

def reverse(nums, start, end)
  while start < end
    nums[start], nums[end] = nums[end], nums[start]
    start += 1
    end -= 1
  end
end

# test
nums = [1, 2, 3]
nextPermutation(nums)
puts nums.join(',')
