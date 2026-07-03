func lengthOfLIS(nums []int) int {
	maxi := -1001
	for _, num := range nums {
		maxi = max(num,maxi)
	}
	cache := make(map[[2]int]int)
    return f(nums, len(nums)-1,maxi+1,cache)
}


func f(nums[]int, ind int, prev int, cache map[[2]int]int) int {
    if ind < 0 {
		return 0
	}
	key := [2]int{ind,prev}
	if _, ok:= cache[key]; ok {
		return cache[key]
	}
	pick, nopick := 0,0
	if nums[ind] < prev {
		pick = 1+f(nums, ind-1, nums[ind], cache)
	}
	nopick = f(nums, ind-1,prev,cache)
	cache[key] = max(pick,nopick)
	return max(pick,nopick)
}