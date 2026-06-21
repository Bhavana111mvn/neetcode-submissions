func merge(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i,j int)bool{
		return intervals[i][0]<intervals[j][0]
	})
	//finalint := make([][]int,len(intervals))
	//var temp []int
	if len(intervals) <=1 {
		return intervals
	}
	finalint := make([][]int, 1, len(intervals))
	finalint[0] = intervals[0]
	for i:=1; i<len(intervals);i++ {
		 temp := finalint[len(finalint)-1]
         if intervals[i][0] <= temp[1] {
			  temp[1] =  maxof(intervals[i][1], temp[1])
		 } else {
			finalint = append(finalint, intervals[i])
		 }
	}
    return finalint
}

func maxof(a,b int) int {
		if a > b{
			return a
		}
		return b
}
