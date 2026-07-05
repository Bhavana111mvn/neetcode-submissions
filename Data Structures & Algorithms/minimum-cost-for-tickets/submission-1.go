func mincostTickets(days []int, costs []int) int {
    dp := make([]int, days[len(days)-1]+30)
    /// recursion ind - 0 -n-1 pass - 0,7,30
    /// dp ind n-1 to 0
    for ind := len(days)-1; ind >=0; ind-- {
        tmp :=make([]int,days[len(days)-1]+30)
        
        for pass:= days[len(days)-1]+29 ; pass >=0; pass-- {
            one := int(1e9)
            if pass >= days[ind] {
                one = 0+dp[pass]
            } else {
                one = costs[0]+dp[0]
            }
            seven := costs[1] + dp[days[ind]+6]
            thirty := costs[2] + dp[days[ind]+29]
            tmp[pass] = min(one,seven,thirty)
 
        }
        dp = tmp
    }
    return dp[0]
}