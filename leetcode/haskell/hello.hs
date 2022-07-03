hello :: String -> IO ()
hello name = putStrLn ("Hello, " ++ name ++ "!")

primes = filterPrime [2..]
  where filterPrime (p:xs) = p : filterPrime [x | x <- xs, x `mod` p /= 0]

# toupper

map toUpper "abc"

map (*99) [1..10]

let (a, b) = (1, 2) in a + b


# pattern matching
let (a:b:c) = "xyz" in a

# ignore value
let (a:_) = "xyz" in a
#math is music
1+2
#charm is love
"Chris"
# put the funk in function
sort [3,1,2]
sort "Chris"
# tuples
(28, "Chris")
fst (28, "Chris")
snd (28, "Chris")
    
# Let them eat cakes
let (a, b) = (1, 2)
let (a, b) = (1, 2)

# construct a list
[1,2,3]
[1..10]
[1,2..10]


