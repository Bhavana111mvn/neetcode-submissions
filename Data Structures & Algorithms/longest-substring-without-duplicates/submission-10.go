func lengthOfLongestSubstring(s string) int {
    strarr := []rune(s)
	hashmap := make(map[rune]int)
    left, right := 0,1
	longlen := 1
	if s== " " {
		return len(s)
	}
	if s == "" {
		return 0
	}
	hashmap[strarr[left]] = 1
	for right < len(strarr) {
		if hashmap[strarr[right]]< 1 {
			hashmap[strarr[right]]=1
			right++
		} else {
		    ////duplicate character
			if longlen < (right-left) {
				longlen = right-left
			}
			for strarr[left] != strarr[right] {
				hashmap[strarr[left]]=0
				left++
			}
            left++
			right++
		}
	}
	if longlen < (right-left) {
				longlen = right-left
	}
    return longlen

}
