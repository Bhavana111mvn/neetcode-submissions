func uniquePaths(m int, n int) int {
    cache := make([][]int, m+1)
    for i:= range cache {
        cache[i]= make([]int, n+1)
        for j:= range cache[i] {
            cache[i][j]=-1
        }
    }
    return f(m-1,n-1,cache)
    
}

func f(i,j int,cache [][]int) int {
    if i==0 && j==0 {
        return 1
    }  
    if cache[i][j] != -1 {
        return cache[i][j]
    }
    up, left := 0,0
    if i-1 >=0 && j>=0 {
        up = f(i-1,j,cache)
    }
    if j-1 >= 0 && i>=0 {
        left = f(i,j-1,cache)
    }
    cache[i][j]= up+left
    return up+left

}