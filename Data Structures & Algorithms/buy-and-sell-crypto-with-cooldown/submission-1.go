func maxProfit(prices []int) int {
    //dp := make(map[[2]int]int)
	dp := make([][]int, len(prices)+2)
	for i:= range dp {
		dp[i]=make([]int,3)
		for j:= range dp[i] {
			dp[i][j]=-1
			if i>=len(prices) {
				dp[i][j]=0
			}
		}
	}
    if len(prices) <=1 {
        return 0
    }
	/// recursion ind 0 to n buy 0-1
	for ind:=len(prices)-1; ind >=0;ind-- {
	//for buy:=1; buy>=0;buy-- {
		for buy:=1; buy>=0;buy--  {
			profit := 0
			if  buy == 0   {
				res := max(-prices[ind] + dp[ind+1][1], dp[ind+1][0])
				profit = max(res,profit)
			}
			////sell case
			if buy==1  {
				profit = max(prices[ind] + dp[ind+2][0],profit,dp[ind+1][1])
			}
			dp[ind][buy] = profit
		}
	}
	fmt.Println(dp)
	return dp[0][0]
   //return f(0,0,prices,dp)  
}

func f(ind, buy int, prices []int ,dp map[[2]int]int) int {
   if ind >= len(prices) {
      return 0
   }
   key := [2]int{ind,buy}
   if _,ok:=dp[key];ok {
      return dp[key]
   }
    profit := 0
        //// buy case
    if  buy == 0   {
        res := max(-prices[ind] + f(ind+1, 1, prices,dp), 0+f(ind+1,0,prices,dp))
        profit = max(res,profit)
    }
        ////sell case
    if buy==1  {
        profit = max(prices[ind] + f(ind+2 , 0,prices,dp),profit,0+f(ind+1,1,prices,dp))
    }
   
   dp[key] = profit
   return profit

}