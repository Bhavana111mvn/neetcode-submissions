func maxProfit(prices []int) int {
	maxi :=-1
	for _, price := range prices {
		maxi = max(maxi, price)
	}
	cache := make([][]int, len(prices))
	for i := range cache {
		cache[i]= make([]int, max(maxi+1,len(prices)))
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
    return f(0,0,prices, cache)
}

func f(ind, buy int, prices []int, cache [][]int) int {
	if ind==len(prices) {
		return 0
	}
	if  cache[ind][buy] != -1 {
		return cache[ind][buy]
	}
	profit := 0
	if buy == 0 {
		profit = max(-prices[ind] + f(ind+1,1,prices, cache),  f(ind+1,0,prices, cache))
	} else {
		profit = max(prices[ind]+ f(ind+1,0,prices, cache),f(ind+1,1,prices, cache))
	}
    cache[ind][buy] = profit
	return profit
}
