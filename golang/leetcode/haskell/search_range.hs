searchRange :: ([Int], Int) -> [Int] 
searchRange (xs, x) = [i | i <- [0..(length xs) - 1], xs !! i == x]

# test

Tes.Hspec.hspec $ do
  describe "searchRange" $ do
    it "should return the index of the first and last occurrence of the element" $ do
      searchRange ([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 1) `shouldBe` [0, 0]
      searchRange ([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 2) `shouldBe` [1, 1]
      searchRange ([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 3) `shouldBe` [2, 2]
      searchRange ([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 4) `shouldBe` [3, 3]
      

