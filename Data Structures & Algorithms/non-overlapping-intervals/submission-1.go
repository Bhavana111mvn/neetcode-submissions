func eraseOverlapIntervals(intervals [][]int) int {
    fmt.Println(intervals)
    sort.Slice(intervals, func(i,j int)bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] < intervals[j][1]
        } 
        return intervals[i][0] < intervals[j][0]
    })
    fmt.Println(intervals)

    prevEnd := intervals[0][1]
    cnt := 0
    for i:=1; i<len(intervals); i++ {
        if intervals[i][0] < prevEnd {
            prevEnd = min(prevEnd, intervals[i][1])
            cnt++
        } else {
           prevEnd = intervals[i][1]
        }

    }
    return cnt
}
