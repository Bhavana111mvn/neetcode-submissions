func findTargetSumWays(nums []int, target int) int {
    return f(len(nums)-1, 0, target, nums)
}

func f(ind, target, sum int, a []int) int {
	if ind == -1 {
		if sum == target {
			return 1
		} else {
			return 0
		}
	}
	plus := f(ind-1, target+a[ind], sum, a)
	minus := f(ind-1, target-a[ind], sum, a)
	//fmt.Println(plus+minus)
	return plus+minus
}

func abs (a int) int {
	if a <0 {
		return -a
	}
	return a
}