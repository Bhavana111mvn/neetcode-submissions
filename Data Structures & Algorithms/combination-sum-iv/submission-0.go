func combinationSum4(nums []int, target int) int {
    dp := make([][]int, target+1)
    for i := range dp {
        dp[i] = make([]int, len(nums)+1)
    }

    // base: 1 way to make 0
    for i := 0; i <= len(nums); i++ {
        dp[0][i] = 1
    }

    for j := 1; j <= target; j++ {
        for i := 1; i <= len(nums); i++ {
            dp[j][i] = dp[j][i-1] // don't use nums[i-1]
            if j >= nums[i-1] {
                dp[j][i] += dp[j-nums[i-1]][len(nums)] // use nums[i-1], and next step can use ALL nums again
            }
        }
    }
    return dp[target][len(nums)]
}
