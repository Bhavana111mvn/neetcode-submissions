func climbStairs(n int) int {
	finalarr := make([]int,2)
	finalarr[0] =1
	finalarr[1] = 2
    if n==1 {
			return 1
	}
	if n==2 {
			return 2
	}
	i :=2
	for i< n {
        tmp := finalarr[1]
		finalarr[1] = finalarr[0] + finalarr[1]
		finalarr[0] = tmp
		i++
	}
	return finalarr[1]
    
}
