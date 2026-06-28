func rob(nums []int) int {
if len(nums) <=1 {
	return nums[0]
}
  dp := make([]int,2)
  dp[0]=0
  dp[1]=nums[1]
  i := 2
  for i<=len(nums)-1 {
	res := max(0+dp[1], nums[i]+dp[0])
	dp[0]=dp[1]
	dp[1]=res
	i++
  }
  res1 :=dp[1]
/// reset dp
  dp[0]=0
  dp[1]=nums[0]
  i = 1
  for i<len(nums)-1 {
	res := max(0+dp[1], nums[i]+dp[0])
	dp[0]=dp[1]
	dp[1]=res
	i++
  }
  return max(dp[1],res1)
    
}
