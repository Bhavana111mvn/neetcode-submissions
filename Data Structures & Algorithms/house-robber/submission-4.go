func rob(nums []int) int {
  dp := make([]int,2)
  dp[0]=0
  dp[1]=nums[0]
  i := 1
  for i<=len(nums)-1 {
	res := max(0+dp[1], nums[i]+dp[0])
	dp[0]=dp[1]
	dp[1]=res
	i++
  }
  return dp[1]
}


