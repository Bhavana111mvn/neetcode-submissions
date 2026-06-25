func coinChange(coins []int, amount int) int {
    cache := make([]int, amount+1)
    res := recurs(amount, amount, coins, cache)
    if res >= 1e9 { return -1 }
    return res
}
func recurs(value, amount int, coins []int, cache []int) int {
    if value == 0 {
        return 0
    } else if value < 0 {
        return 1e9
    }
    if cache[value] !=0 {
        return cache[value]
    }
    tmp := int(1e9)
    for i:=0; i<len(coins);i++ {
        res := 1+recurs(value-coins[i], amount, coins, cache)
        tmp = min(res,int(tmp))
        if value-coins[i] > 0 {
           cache[value-coins[i]] = res-1
        }
    }

    return int(tmp)
}
