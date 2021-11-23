LongestValidParentheses: String -> Int
LongestValidParentheses str =
  let
    (str', _) = foldl (\(str, last) c ->
      if c == '('
        then (str ++ [c], c)
        else if last == '('
          then (str, last)
          else (str ++ [c], c))
      (str, ' ') str
    ) ("", ' ') str
  in
    length str'

-- Test

main :: IO ()
main = do
  print $ longestValidParentheses "()"
  print $ longestValidParentheses "()()"
  print $ longestValidParentheses "()(()"
  print $ longestValidParentheses "()(()()"
  print $ longestValidParentheses "()(()())"
  print $ longestValidParentheses "()(()()))"
  print $ longestValidParentheses "()(()())))"
