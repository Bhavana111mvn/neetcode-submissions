func productExceptSelf(nums []int) []int {
   m := len(nums)
   right :=make([]int,m)
   left :=make([]int,m)
   left[0]=1
   right[m-1]=1
   product:=make([]int,m)
   product[0]=1
   for i:=1; i<len(nums);i++ {
	   left[i] = left[i-1]*(nums[i-1])
	   product[i] = left[i]
   }
   for i:=len(nums)-2; i>=0;i-- {
	 right[i] = right[i+1]*(nums[i+1])
	  product[i] *=right[i]
   }
   return product


}
