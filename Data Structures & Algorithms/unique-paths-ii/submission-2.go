func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for ind, _ := range dp {
		dp[ind] = make([]int, n)
	}
	dp[0][0] = 0
	if obstacleGrid[0][0] !=1 {
		dp[0][0]=1
	}
	for i:=0;i<m;i++ {
		for j:=0;j<n;j++ {
			if i==0 && j==0 {
				continue
			}
			if i-1 <0 && j-1 <0 || i>=0 && j>=0 && obstacleGrid[i][j]==1 {
				dp[i][j] = 0
				continue
			}
			if i-1 < 0 {
				dp[i][j]=dp[i][j-1]
			} else if j-1 <0 {
				dp[i][j]=dp[i-1][j]
			} else {
				dp[i][j]=dp[i-1][j] + dp[i][j-1]
			}
			fmt.Println(dp[i][j])
		}
	}

	return dp[m-1][n-1]
}
