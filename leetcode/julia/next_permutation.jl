function next_permutation(nums::Vector{Int})
    next_perm!(nums)
    return nums
end

function next_perm!(itr)::Bool                      
    (isempty(itr) || length(itr) == 1) && return false
        
    i = findlast(idx -> itr[idx + 1] > itr[idx], 1:lastindex(itr) - 1)             
    isnothing(i) && (reverse!(itr); return false)
      
    j = findlast(idx -> itr[idx] > itr[i], i+1:lastindex(itr)) + i
    itr[i], itr[j] = itr[j], itr[i] 
    reverse!(itr, i + 1)
    return true
end

# test
print("[ ")
for i in next_permutation([1,2,3])
	print(i)
  print(" ")
end
print("]")
