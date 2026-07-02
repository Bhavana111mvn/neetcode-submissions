func canPartition(nums []int) bool {
    total :=0
    for _, num := range nums {
        total+=num
    }
    if total%2 !=0 {
        return false
    }
    dp := make([][]bool, len(nums)+1)
    for i := range dp {
        dp[i] = make([]bool, total)
    }
    dp[0][nums[0]] = true
    ind,target :=0,0
    for ind=1; ind<len(nums);ind++ {
        dp[ind][0] = true
        take, notake := false,false
        for target=0;target<=total/2; target++ {
            if nums[ind] <=target {
                 take = dp[ind-1][target-nums[ind]]
            }
            notake = dp[ind-1][target]
            dp[ind][target] = take|| notake
            //fmt.Println(dp[ind][target])
        }
    }
    return dp[ind-1][target-1]
    //return f(len(nums)-1,total/2,nums,dp)
}

// func f(ind, target int, a[]int, dp [][]bool) bool {
//     if target == 0 {
//         return true
//     }
//     if ind==0 && target==a[ind] {
//         return true
//     } else if ind==0 {
//         return false
//     }
//     if dp[ind][target] {
//         return true
//     }
//     take, notake := false,false
//     if a[ind] <=target {
//         take = f(ind-1,target-a[ind],a,dp)
//     }
//     notake = f(ind-1,target,a,dp)
//     dp[ind][target] = take|| notake
//     return take|| notake
// }
