func maxProfit(prices []int) int {
	maxi :=-1
	for _, price := range prices {
		maxi = max(maxi, price)
	}
	cache := make([][]int, len(prices))
	for i := range cache {
		cache[i]= make([]int, maxi+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
    return f(0,-1,prices, cache)
}

func f(ind, buy int, prices []int, cache [][]int) int {
	if ind==len(prices) {
		return 0
	}
	if buy != -1 && cache[ind][buy] != -1 {
		return cache[ind][buy]
	}
	optbuy, optsell, nopick := 0,0,0
	if buy == -1 {
		optbuy = 0 + f(ind+1,prices[ind],prices, cache)
	} else {
		optsell = prices[ind]-buy + f(ind+1,-1,prices, cache)
	}
	nopick = 0 + f(ind+1,buy,prices, cache)
	if buy != -1 {
       cache[ind][buy] = max(optbuy,optsell,nopick)
	}
	return max(optbuy,optsell,nopick)
}
