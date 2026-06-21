func missingNumber(nums []int) int {
	sum := 0
	total := len(nums)*(len(nums)+1)/2
	for _, num := range nums {
		sum+=num
	}
	return total-sum

}
