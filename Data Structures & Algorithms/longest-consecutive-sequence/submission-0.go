func longestConsecutive(nums []int) int {
   hashmap := make(map[int]bool)
   for _, num := range nums {
	 hashmap[num]=true
   }
   longlen :=0
   for i:=0; i<len(nums);i++ {
	    start := nums[i]
	    if !hashmap[start-1] {
			/// start is fine 
			j :=1
			for hashmap[start+j] {
				j++
			}
			if longlen < j {
				longlen = j
			}
		}
   }
   return longlen
}
