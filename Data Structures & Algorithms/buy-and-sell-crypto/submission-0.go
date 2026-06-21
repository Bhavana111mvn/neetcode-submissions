func maxProfit(prices []int) int {
    left, right := 0, 1
    profit :=0
    for left < len(prices)-1 {
        if profit < prices[right] - prices[left] {
            profit = prices[right] - prices[left]
        }
        right++
        if right == len(prices) {
            /// reset
            left++
            right = left+1
        }
    }
    return profit

}
