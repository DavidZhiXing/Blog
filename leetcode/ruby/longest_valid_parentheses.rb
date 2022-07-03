def longestValidParentheses(s)
  return 0 if s.length < 2
  stack = []
  max = 0
  s.each_char.with_index do |c, i|
    if c == '('
      stack.push(i)
    else
      if stack.length > 0
        stack.pop
        max = [max, i - stack.last].max
      else
        stack.push(i)
      end
    end
  end
  max
end

# test
puts longestValidParentheses('(()')
puts longestValidParentheses(')()())')
puts longestValidParentheses(')()())')