func maxSubArray(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    sum, maxsum := nums[0], nums[0]

    for right := 1; right < len(nums); right++ {
        if sum < 0 {
            sum = nums[right] // start new subarray
        } else {
            sum += nums[right] // extend existing subarray
        }

        if sum > maxsum {
            maxsum = sum
        }
    }

    return maxsum
}