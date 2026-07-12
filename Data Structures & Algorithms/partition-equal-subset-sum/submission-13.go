func canPartition(nums []int) bool {
    total :=0
    for _, num := range nums {
        total+=num
    }
    if total%2 !=0 {
        return false
    }
    target := total/2
    dp :=make([]bool, target+1)
    dp[0] = true
    for i:=0;i<len(nums);i++ {
        tmp := make([]bool, target+1)
        tmp[0] = true
        for j:=1; j<target+1;j++ {
            if j >= nums[i] { // <-- add this
                tmp[j] = dp[j-nums[i]] || dp[j]
            } else {
                tmp[j] = dp[j]
            }
        }
        dp = tmp
    }
    return dp[target]
}


