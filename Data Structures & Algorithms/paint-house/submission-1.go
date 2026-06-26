func minCost(costs [][]int) int {
	cache := make([][3]int, len(costs))
    return min(dp(costs,0,0,cache),dp(costs,0,1,cache),dp(costs,0,2,cache))
}


func dp(costs [][]int, i int, j int,cache [][3]int) int {
	if i >= len(costs) {
		return 0
	}
	if 0!=cache[i][j] {
		return cache[i][j]
	}
    res := int(1e9)
    for k:=0;k<3;k++ {
		if k!=j {
			res = min(dp(costs,i+1,k,cache),res)
		}
	}
	cache[i][j]= costs[i][j] + res
    return costs[i][j] + res
}