def search_in_rotated_sorted_list(nums, target)
  return -1 if nums.nil? || nums.empty?
  return 0 if nums.size == 1 && nums[0] == target
  return -1 if nums.size == 1 && nums[0] != target
  left = 0
  right = nums.size - 1
  while left <= right
    mid = (left + right) / 2
    if nums[mid] == target
      return mid
    elsif nums[mid] < target
      if nums[left] == target
        return left
      elsif nums[left] < target
        left = mid + 1
      else
        right = mid - 1
      end
    else
      if nums[right] == target
        return right
      elsif nums[right] > target
        right = mid - 1
      else
        left = mid + 1
      end
    end
  end
  return -1
end

// test
nums = [4, 5, 6, 7, 0, 1, 2]
target = 0
puts "nums: #{nums}, target: #{target}"
puts "result: #{search_in_rotated_sorted_list(nums, target)}"
