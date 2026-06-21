func canJump(nums []int) bool {
	left:= len(nums)-2
	maxsum := make([]int, len(nums))
	if len(nums) == 1 {
		return true
	}
	if nums[0]==0 {
		return false
	}
	for left >0 {
		i:=1
		for i <= nums[left] {
			if left+i < len(nums) && ((i+ maxsum[left+i]) > maxsum[left]) {
                maxsum[left]=i+ maxsum[left+i]
			}
			i++
		}
		left--	
	}
	for i:=1; i<=nums[0]; i++ {
		if i+maxsum[i] >= len(nums)-1 {
			return true
		}
	}
    return false
}
