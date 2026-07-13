func change(amount int, coins []int) int {
	dp := make([][]int, len(coins)+1)
	for i:= range dp {
		dp[i]=make([]int, amount+1)
	}
	for i:=0; i<len(dp);i++ {
		dp[i][0]= 1
	}
	dp[0][0]= 1
	for i:=1; i<len(dp);i++ {
		for j:=1; j<len(dp[0]); j++ {
			if j>=coins[i-1] {
				dp[i][j]=dp[i][j-coins[i-1]] + dp[i-1][j]
			} else {
                dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(coins)][amount]
}
