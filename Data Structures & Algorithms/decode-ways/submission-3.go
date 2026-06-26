func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	cache := make(map[int]int) // key = index, not string
	return dfs(s, 0, cache)
}

func dfs(s string, index int, cache map[int]int) int {
	if index == len(s) {
		return 1
	}
	if s[index] == '0' {
		return 0
	}
	
	// Check cache
	if val, ok := cache[index]; ok {
		return val
	}
	
	ways := 0
	
	// 1 digit
	ways += dfs(s, index+1, cache)
	
	// 2 digits
	if index+1 < len(s) {
		num := int(s[index]-'0')*10 + int(s[index+1]-'0')
		if num >= 10 && num <= 26 {
			ways += dfs(s, index+2, cache)
	}
	}
	
	cache[index] = ways // store result for this index
	return ways
}