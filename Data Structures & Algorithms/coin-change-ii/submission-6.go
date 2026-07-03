func change(amount int, coins []int) int {
	dp := make([][]int, len(coins)+1)
	for i := range dp {
		dp[i] = make([]int, amount+1) // was len(coins)+1
		dp[i][0] = 1
	}
	for target:=0; target<=amount;target++  {
		if target %coins[0] ==0 {
			dp[0][target]= 1
		}
	}
	for ind:=1; ind<len(coins);ind++ {
		for target:=0; target<=amount;target++ {
			left, right:= 0,0
			if target-coins[ind] >= 0 {
               left = dp[ind][target-coins[ind]]
			}
			right = dp[ind-1][target]			
			dp[ind][target] = left + right
		}
	}
	//fmt.Println(dp)
	return dp[len(coins)-1][amount]
}
