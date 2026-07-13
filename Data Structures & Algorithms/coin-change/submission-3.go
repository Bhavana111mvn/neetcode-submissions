func coinChange(coins []int, amount int) int {
    dp :=make([][]int, len(coins)+1)
    for i:= range dp {
        dp[i] = make([]int, amount+1)
    }
    for i:=0; i<len(dp[0]);i++ {
        dp[0][i]= int(1e9)
    }
    for i:=1; i<len(dp);i++ {
        for j:=1; j<len(dp[0]); j++ {
            if j>=coins[i-1] {
                dp[i][j]=min(1+dp[i][j-coins[i-1]], dp[i-1][j])
            } else {
                dp[i][j] = dp[i-1][j]
            }
        }
    }
    if dp[len(coins)][amount] == int(1e9) {
        return -1
    }
    return dp[len(coins)][amount]
}

