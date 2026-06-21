func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
    cnt := 0
	left, right :=0, len(people)-1
	for left <= right {
		if people[right]+people[left] <= limit {
			left++
		}
		cnt++
		right--
	}
	return cnt
}
