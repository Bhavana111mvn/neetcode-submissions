func findTargetSumWays(nums []int, target int) int {
	dp := make(map[[2]int]int)
    return f(len(nums)-1, 0, target, nums,dp)
}

func f(ind, target, sum int, a[]int,dp map[[2]int]int) int {
	if ind == -1 {
		if sum == target {
			return 1
		} else {
			return 0
		}
	}
	key := [2]int{ind, target}
    if _, ok:= dp[key]; ok {
		return dp[key]
	}
	plus := f(ind-1, target+a[ind], sum, a,dp)
	minus := f(ind-1, target-a[ind], sum, a,dp)
	dp[key] = plus+minus
	//fmt.Println(plus+minus)
	return plus+minus
}

func abs (a int) int {
	if a <0 {
		return -a
	}
	return a
}