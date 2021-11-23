searchInRotatedSortedArray: [int] -> int
searchInRotatedSortedArray nums = searchInRotatedSortedArray' nums 0 (length nums - 1)
searchInRotatedSortedArray' nums start end
    | start > end = -1
    | nums !! mid == target = mid
    | nums !! mid < target = searchInRotatedSortedArray' nums (mid + 1) end
    | otherwise = searchInRotatedSortedArray' nums start (mid - 1)
    where mid = (start + end) `div` 2

    