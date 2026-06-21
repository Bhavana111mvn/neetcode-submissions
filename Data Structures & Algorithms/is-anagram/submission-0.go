func isAnagram(s string, t string) bool {
	hashmap := make(map[rune]int)
	for _, val := range s{
		hashmap[val]++
	}
	for _, val := range t {
		hashmap[val]--
	}
	for _, val := range hashmap {
		if val != 0 {
			return false
		}
	}
    return true
}
