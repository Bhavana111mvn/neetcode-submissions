func numSquares(n int) int {
	squares := []int{}
	for i := 1; i*i <= n; i++ {
		squares = append(squares, i*i)
	}

	dp := make([][]int, len(squares)+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// base: impossible = inf, 0 sum = 0 squares
	for i := 0; i <= len(squares); i++ {
		dp[i][0] = 0
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = 100000 // inf
	}

	for i := 1; i <= len(squares); i++ {
		for j := 1; j <= n; j++ {
			if j >= squares[i-1] {
				// take it: dp[i][j-squares[i-1]] + 1, or skip it: dp[i-1][j]
				dp[i][j] = min(dp[i][j-squares[i-1]]+1, dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(squares)][n]
}

