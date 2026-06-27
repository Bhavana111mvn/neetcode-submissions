func minExtraChar(s string, dictionary []string) int {
    hashmap := make(map[string]bool)
    for _, s := range dictionary {
        hashmap[s] = true
    }
    cache := make([]int, len(s))
    return dfs(0, s, hashmap, cache)
}

func dfs(i int, s string, hashset map[string]bool, cache []int) int {
    if i == len(s) {
        return 0
    }
    if cache[i] !=0 {
        return cache[i]
    }

    res := 1 + dfs(i+1,s,hashset,cache)
    for j:=i; j<len(s); j++ {
        if hashset[s[i:j+1]] {
            res = min(dfs(j+1,s, hashset,cache), res)
        }
    }
    cache[i] = res
    return res

}