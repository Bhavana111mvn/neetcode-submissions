func shortestWay(source string, target string) int {
	pos := make(map[byte]bool)
	for ind := 0; ind < len(source); ind++ {
		pos[source[ind]] = true
	}
	left, cnt := 0, 0
	for left < len(target) {
		if !pos[target[left]] {
			return -1
		}
		sInd := 0
		cnt++
		for left < len(target) && sInd < len(source) {
			if target[left] == source[sInd] {
				left++
			}
			sInd++
		}
	}
	return cnt
}