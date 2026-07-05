func findMaxForm(strs []string, m int, n int) int {
    dp := make([][]int, m+1)
    for i:= range dp {
        dp[i]= make([]int,n+1)
        for j:= range dp[i] {
            dp[i][j]=0
        }
    }
    //// recursion : ind last -0  m,n- final to 0
    //zero, one := m,n
    //ind := 0
    for ind := 0; ind <len(strs);ind++ {
        tmp := make([][]int, m+1)
        for i:= range tmp {
            tmp[i]= make([]int,n+1)
        }
        count0, count1 := 0,0
        for _,val := range strs[ind] {
            if val=='1' {
                count1++
            } else {
               count0++
            }
        }
        for zero:=0; zero <=m;zero++ {
            for one:=0; one<=n;one++ {
                pick := 0
                nopick := 0+dp[zero][one]
                if count0<=zero && count1<=one {
                    pick = 1+dp[zero-count0][one-count1]
                }
                
                tmp[zero][one] = max(pick,nopick)

            }
        
        }
        dp = tmp
    }
    //fmt.Println(dp)
    return dp[m][n]
}
