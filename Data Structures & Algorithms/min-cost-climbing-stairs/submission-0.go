func minCostClimbingStairs(cost []int) int {
	dp := make([]int,len(cost)+1)
	dp[0] = 0
	dp[1] = 0
	i := 0
	for i=2 ; i<=len(cost); i++ {
		dp[i] = min(cost[i-1]+dp[1], cost[i-2]+dp[0])
		dp[0] = dp[1]
		dp[1] = dp[i]
	}
	return dp[1]
    
}
