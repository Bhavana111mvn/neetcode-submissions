func maxProfit(prices []int) int {
    max := prices[len(prices)-1]
    profit :=0
    left := len(prices)-2 
    for left >= 0 {
        if prices[left] > max {
            max = prices[left]
        }
        if max-prices[left] > profit {
            profit = max-prices[left]
        }
        left--
    }

    return profit

}
