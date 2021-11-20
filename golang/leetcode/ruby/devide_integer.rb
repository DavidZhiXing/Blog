def devideTwoInteger num1, num2
  return 0 if num2 == 0
  return num1 if num2 == 1
  return -num1 if num2 == -1
  return 0 if num1 == 0
  return 0 if num1 == -1 && num2 == -1
  return -1 if num1 == -1 && num2 == 1

  result = 0
  sign = 1
  if num1 < 0
    sign = -1
    num1 = -num1
  end
  if num2 < 0
    sign = -sign
    num2 = -num2
  end
  while num1 >= num2
    num1 -= num2
    result += 1
  end
  return result * sign
end

// test
puts devideTwoInteger(10, 3)
puts devideTwoInteger(10, -3)
puts devideTwoInteger(-10, 3)
puts devideTwoInteger(-10, -3)
puts devideTwoInteger(10, 0)
puts devideTwoInteger(0, 10)
puts devideTwoInteger(0, 0)
puts devideTwoInteger(0, -10)
puts devideTwoInteger(-10, 0)
puts devideTwoInteger(-10, -10)
puts devideTwoInteger(-10, -1)
puts devideTwoInteger(-1, -10)