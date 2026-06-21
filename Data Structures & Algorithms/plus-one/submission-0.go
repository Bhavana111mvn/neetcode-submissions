func plusOne(digits []int) []int {
	val := 0
    for _, digit := range digits {
        val = val*10 + digit
	}
	val++
	fmt.Println(val)
	finalarr := make([]int,0)
	for val > 0 {
		x := val %10
		val = val/10
		finalarr = append(finalarr,x)
	}
	fmt.Println(finalarr)
	for ind :=0; ind < len(finalarr)/2; ind++ {
		tmp := finalarr[ind]
		finalarr[ind] = finalarr[len(finalarr)-1-ind]
		finalarr[len(finalarr)-1-ind] = tmp
	} 
	return finalarr
}
