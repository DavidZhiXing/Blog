def searchRange(nums, target)
  return [-1, -1] if nums.empty?
  left, right = 0, nums.size - 1
  while left <= right
    mid = (left + right) / 2
    if nums[mid] == target
      return [mid, mid]
    elsif nums[mid] < target
      if nums[left] == target
        return [left, mid]
      elsif nums[left] < target
        left = mid + 1
      else
        right = mid - 1
      end
    else
      if nums[right] == target
        return [mid, right]
      elsif nums[right] > target
        right = mid - 1
      else
        left = mid + 1
      end
    end
  end
  [-1, -1]
end

# Test
nums = [5, 7, 7, 8, 8, 10]
target = 8
puts searchRange(nums, target)
