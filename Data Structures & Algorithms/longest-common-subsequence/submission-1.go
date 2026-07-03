func longestCommonSubsequence(text1 string, text2 string) int {
    cache := make([][]int, len(text1))
    for i:= range cache {
        cache[i] = make([]int, len(text2))
        for j:= range cache[i] {
            cache[i][j]=-1
        }
    }
    return f(text1,text2,0,0,cache)
}


func f(text1 string, text2 string, text1Id int, text2Id int,cache [][]int) int {
    if text1Id >= len(text1) || text2Id >= len(text2) {
        return 0
    }
    if cache[text1Id][text2Id] != -1 {
        return cache[text1Id][text2Id]
    }
    case1, case2, case3 := 0,0,0
    if text1[text1Id] == text2[text2Id] {
        case1 = 1+f(text1,text2,text1Id+1,text2Id+1,cache) 
    } 
    case2 = 0 +f(text1,text2,text1Id, text2Id+1,cache)
    case3 = 0 +f(text1,text2,text1Id+1, text2Id,cache)
    cache[text1Id][text2Id]= max(case1, case2, case3)
    return max(case1, case2, case3)

}