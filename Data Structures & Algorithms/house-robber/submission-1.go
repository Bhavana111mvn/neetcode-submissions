func rob(nums []int) int {
  dp := make([]int,2)
  dp[0]=0
  dp[1]=0
  i := len(nums)-1
  for i>=0 {
	  tmp := dp[1]
	  dp[1]= dp[0]
	  out := max(nums[i]+tmp, dp[0])
	  dp[0]= out
	  i--
  }
  return dp[0]
}


func max(a,b int)int {
	if a>b {
		return a
	}
    return b
}
