func minCostII(costs [][]int) int {
    if len(costs) == 0 { return 0 }
    cache := make([][]int, len(costs))
    for i := range cache {
        cache[i] = make([]int, len(costs[0]))
    }
    res := int(1e9)
    for i := 0; i < len(costs[0]); i++ {
        res = min(dp(costs, 0, i, cache), res)
    }
    return res

}

func dp(costs [][]int, i int, j int,cache [][]int) int {
	if i >= len(costs) {
		return 0
	}
	if 0!=cache[i][j] {
		return cache[i][j]
	}
    res := int(1e9)
    for k:=0;k<len(costs[0]);k++ {
		if k!=j {
			res = min(dp(costs,i+1,k,cache),res)
		}
	}
	cache[i][j]= costs[i][j] + res
    return costs[i][j] + res
}