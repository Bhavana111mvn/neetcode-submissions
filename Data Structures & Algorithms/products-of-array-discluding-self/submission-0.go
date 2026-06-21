func productExceptSelf(nums []int) []int {
   m := len(nums)
   right :=make([]int,m)
   left :=make([]int,m)
   left[0]=1
   right[m-1]=1
   product:=make([]int,m)
   for i:=1; i<len(nums);i++ {
	   left[i] = left[i-1]*(nums[i-1])
   }
   for i:=len(nums)-2; i>=0;i-- {
	 right[i] = right[i+1]*(nums[i+1])
   }

   for i:=0; i<len(nums);i++ {
         product[i] = left[i]*right[i]
   }
   return product


}
