def strStr(haystack, needle)
  return 0 if needle.length == 0
  return -1 if haystack.length == 0
  return -1 if needle.length > haystack.length
  return haystack.index(needle)
end

// test
puts strStr("hello", "ll")
puts strStr("aaaaa", "bba")
puts strStr("aaaaa", "")
puts strStr("", "")
puts strStr("", "a")
puts strStr("a", "")
puts strStr("a", "a")