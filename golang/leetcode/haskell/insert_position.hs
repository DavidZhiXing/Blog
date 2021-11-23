insertPosition :: (Ord a) => a -> [a] -> Int
insertPosition x [] = 0
insertPosition x (y:ys)
    | x < y = 0
    | otherwise = 1 + insertPosition x ys

""" Test Cases """
Tes.Hspec.Spec.describe "insertPosition" $ do
    it "should return 0 for an empty list" $ do
        insertPosition 1 [] `shouldBe` 0
    it "should return 0 for a list with one element" $ do
        insertPosition 1 [2] `shouldBe` 0
    it "should return 0 for a list with two elements" $ do
        insertPosition 1 [2, 3] `shouldBe` 0
    it "should return 1 for a list with three elements" $ do
        insertPosition 3 [2, 3, 4] `shouldBe` 1

