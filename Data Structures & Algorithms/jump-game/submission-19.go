func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	maxsum := make([]int, len(nums))
	
	for left := len(nums) - 2; left >= 0; left-- { // >= 0 instead of > 0
		for i := 1; i <= nums[left]; i++ {
			if left+i < len(nums) && i+maxsum[left+i] > maxsum[left] {
				maxsum[left] = i + maxsum[left+i]
			}
	}
	}
	return maxsum[0] >= len(nums)-1 // single return
}