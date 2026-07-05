func maxProfit(prices []int) int {
    dp := make(map[[2]int]int)
    if len(prices) <=1 {
        return 0
    }
   return f(0,0,prices,dp)  
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