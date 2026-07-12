func canPartition(nums []int) bool {
    total :=0
    for _, num := range nums {
        total+=num
    }
    if total%2 !=0 {
        return false
    }
    target := total/2
    dp := make([][]bool, len(nums)+1)
    for i:= range dp {
        dp[i]=make([]bool, target+1)
    }
    for i := 0; i <= len(nums); i++ {
        dp[i][0] = true // sum 0 is always possible
    }
    for i:=1;i<len(dp);i++ {
        for j:=1; j<len(dp[0]);j++ {
            if j >= nums[i-1] { // <-- add this
                dp[i][j] = dp[i-1][j-nums[i-1]] || dp[i-1][j]
            } else {
                dp[i][j] = dp[i-1][j]
            }
        }
    }
    return dp[len(dp)-1][len(dp[0])-1]
}


