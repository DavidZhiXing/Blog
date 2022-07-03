strStr : String -> String -> String
strStr s1 s2 =
  let
    s1' = T.pack s1
    s2' = T.pack s2
  in
    T.unpack $ T.drop (T.length s2') $ T.breakOn s2' s1'


