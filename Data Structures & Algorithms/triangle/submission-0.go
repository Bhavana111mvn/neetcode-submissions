func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	n := len(triangle[m-1])
	dp := make([]int, n)
	tmp := make([]int,n)
	for ind, _ := range dp {
		dp[ind] = triangle[m-1][ind]
		tmp[ind] = dp[ind]
	}
	for i:=m-2; i>=0;i-- {
		for j:=0;j<len(triangle[i]);j++ {
			tmp[j] = triangle[i][j]+ min(dp[j],dp[j+1])
		}
		dp = tmp
	}
	return dp[0]
}
