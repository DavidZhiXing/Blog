nextPermutation: [int] -> [int]
nextPermutation xs = nextPermutation' xs (length xs - 1)

nextPermutation' : [int] -> int -> [int]
nextPermutation' [] _ = []
nextPermutation' xs 0 = xs
nextPermutation' xs i =
  let
    x = xs !! i
    xs' = take i xs ++ drop (i + 1) xs
    xs'' = xs' ++ [x]
  in
    if xs' == xs
    then xs
    else nextPermutation' xs'' (i - 1)