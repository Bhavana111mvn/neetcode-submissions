func maxProfit(prices []int) int {
	// cache := make([][]int, len(prices)+1)
	// for i := range cache {
	// 	cache[i]= make([]int, 2)
	// 	for j := range cache[i] {
	// 		cache[i][j] = -1
	// 		if i == len(prices) {
	// 			cache[i][j]=0
	// 		}
	// 	}
	// }
	cache := make([]int, 2)
	for ind:= len(prices)-1; ind >=0;ind-- {
		profit := 0
		for buy:=1; buy >=0;buy-- {
			if buy == 0 {
				profit = max(-prices[ind] + cache[1],  cache[0])
		       //profit = max(-prices[ind] + cache[ind+1][1],  cache[ind+1][0])
	        } else {
				profit = max(prices[ind] + cache[0],  cache[1])
		        //profit = max(prices[ind]+ cache[ind+1][0], cache[ind+1][1])
	        }
			//cache[ind][buy] = profit
			cache[buy] = profit
			//fmt.Println(cache)
		}
	}
    return cache[0]
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
