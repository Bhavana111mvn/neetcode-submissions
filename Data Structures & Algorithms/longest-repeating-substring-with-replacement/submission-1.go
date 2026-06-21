func characterReplacement(s string, k int) int {
	maxcount, res, left, right := 0,0,0,0
	count := make(map[byte]int)
	for right=0; right< len(s); right++ {
        count[s[right]]++
		if maxcount < count[s[right]] {
			maxcount = count[s[right]]
		}

		if ((right-left+1 -maxcount) > k) {
			    count[s[left]]--
				left++
		}

		if (right-left+1) > res {
			res = right-left+1
		}


	}

   return res
}
