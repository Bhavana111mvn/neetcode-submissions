func checkInclusion(s1 string, s2 string) bool {
	freq := make(map[byte]int)
	for ind:=0; ind<len(s1); ind++ {
		freq[s1[ind]]++
	}
    left,cnt := 0,0
	tmp := 0
	start := false
	for left < len(s2) {
		if freq[s2[left]] > 0 {
		   freq[s2[left]]--
		   if cnt==0 && !start {
			  start = true
			  tmp = left
		   }
		   cnt++
		   if start && cnt == len(s1) {
			  return true
		   }
		   left++
		} else {
			if start {
				start = false
			    cnt = 0
				i:= tmp
				for i< left {
					freq[s2[i]]++
					i++
				}
				left = tmp+1
			} else {
				left++
			}
		}

	}
    return false
}
