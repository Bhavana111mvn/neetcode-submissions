func minPathSum(grid [][]int) int {
   m,n := len(grid), len(grid[0])
   dp := make([][]int, m)
   for ind, _ := range dp {
	  dp[ind] = make([]int, n)
   }

    for i:=0;i<m;i++ {
	  for j:=0; j<n;j++ {
		  if i==0 && j==0 {
			dp[0][0] = grid[0][0]
			continue
		  }
		  left, right := int(1e9),int(1e9)
		  if i-1>=0 {
            left = dp[i-1][j]+grid[i][j]
		  }
		  if j-1 >=0 {
			right = dp[i][j-1]+grid[i][j]
			//fmt.Println(right)
		  }
		  dp[i][j]= min(left,right)
		  if dp[i][j]== int(1e9) {
			dp[i][j]=0// reset
		  }
	   }
    }
    return dp[m-1][n-1]
	
}
