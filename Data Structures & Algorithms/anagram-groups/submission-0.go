func groupAnagrams(strs []string) [][]string {
    hashmap := make(map[[26]int][]string)
	finalstr := make([][]string,0)
    for _, str := range strs {
		var engarr [26]int
		for _, ch := range str {
			engarr[ch-'a']++
		}
		hashmap[engarr]= append(hashmap[engarr],str)

	}
	for _, lst := range hashmap {
		finalstr = append(finalstr, lst)
	}
    return finalstr
}
