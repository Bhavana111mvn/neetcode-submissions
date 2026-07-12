func findMaxForm(strs []string, m int, n int) int {
    dp := make([][]int, m+1)
	for i:= range dp {
		dp[i]=make([]int, n+1)
		for j:= range dp[i] {
			dp[i][j]=0
		}
	}
	dp[0][0] = 0
	for k:= 0; k<len(strs);k++ {
		count0,count1 := 0,0
		for _, val := range strs[k] {
			if val == '0' {
				count0++
			} else {
				count1++
			}
		}
		tmp := make([][]int, m+1)
		for i:= range tmp {
			tmp[i]=make([]int, n+1)
		}
		for i:=0; i<=m;i++ {
			for j:=0; j<=n;j++ {
				if i>= count0 && j>= count1 {
					tmp[i][j] = max(1+dp[i-count0][j-count1], dp[i][j])
				} else {
					tmp[i][j]=dp[i][j]
				}
			}
		}
		dp = tmp
	}
	return dp[m][n]
}
