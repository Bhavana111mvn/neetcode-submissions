func pivotIndex(nums []int) int {
    prefix ,suffix := 0,0
    for _, num := range nums {
         suffix+=num
    }
    for ind, num := range nums {
        if ind !=0 {
            prefix += nums[ind-1]
        }
        suffix -=num
        if prefix == suffix {
            return ind
        }

    }
    return -1
	
}
