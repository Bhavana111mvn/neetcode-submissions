func lengthOfLIS(nums []int) int {
	cache := make([][]int, len(nums)+1)
	for i:= range cache {
		cache[i]= make([]int, len(nums)+1)
		for j:= range cache[i] {
			cache[i][j]=-1
		}
	}
    return f(0,-1,nums,cache)
}


func f(ind int, prev_ind int, nums[]int,cache [][]int) int {
	if ind == len(nums) {
		return 0
	}
	if prev_ind != -1 && cache[ind][prev_ind] != -1 {
        return cache[ind][prev_ind]
	}
	pick := 0
	nopick := 0+f(ind+1,prev_ind,nums,cache)
	if (prev_ind == -1 || nums[ind] >nums[prev_ind]) {
       pick = 1 + f(ind+1,ind,nums,cache)
	}
	if prev_ind != -1 {
		cache[ind][prev_ind]= max(pick,nopick)
	}
	return max(pick,nopick)
}