func subarraySum(nums []int, k int) int {
   res, currsum := 0,0
   prefixsum := map[int]int{0:1}

   for _, num := range nums {
	    currsum += num
		diff := currsum -k
		res += prefixsum[diff]
		prefixsum[currsum]++
   }
   return res
}
