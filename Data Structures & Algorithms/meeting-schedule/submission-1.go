/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
    sort.Slice(intervals, func(i,j int)bool {
        return intervals[i].start < intervals[j].start
    })
    if len(intervals) ==0 {
        return true
    }
    prev := intervals[0]
    for i:=1; i<len(intervals);i++ {
        if intervals[i].start < prev.end {
            return false
        }
        prev = intervals[i]
    }
    return true
}
