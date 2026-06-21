func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	finalarr := make([]int,0)
	for ind, num := range nums {
		if val, ok :=hashmap[target-num]; ok  && ind != val {
			finalarr = append(finalarr, val)
			finalarr = append(finalarr, ind)
			break
		}
	    hashmap[num] = ind
	}
	return finalarr
    
}
