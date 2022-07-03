def insertPosition(nums, target)
  nums.each_with_index do |num, i|
    return i if num > target
  end
  nums.length
end

# test
nums = [1,3,5,6]
target = 5
puts insertPosition(nums, target)
