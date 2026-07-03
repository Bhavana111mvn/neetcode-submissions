func longestCommonSubsequence(text1 string, text2 string) int {
    cache := make([][]int, len(text1)+1)
    for i:= range cache {
        cache[i] = make([]int, len(text2)+1)
        for j:= range cache[i] {
            cache[i][j]=0
        }
    }
    // Initialize (0,0)
    if text1[0] == text2[0] {
        cache[0][0] = 1
    }

    // Initialize first row
    for j := 1; j < len(text2); j++ {
        if text1[0] == text2[j] {
            cache[0][j] = 1
        } else {
            cache[0][j] = cache[0][j-1]
        }
    }

    // Initialize first column
    for i := 1; i < len(text1); i++ {
        if text1[i] == text2[0] {
            cache[i][0] = 1
        } else {
            cache[i][0] = cache[i-1][0]
        }
    }
    for text1Id:= 1; text1Id<len(text1);text1Id++ {
        for text2Id:=1; text2Id<len(text2);text2Id++ {
            case1, case2, case3 := 0,0,0
            if text1[text1Id] == text2[text2Id]  {
                case1 = 1+cache[text1Id-1][text2Id-1]
            } 
            if text2Id-1 >= 0 {
               case2 = 0 +cache[text1Id][text2Id-1] 
            }
            case3 = 0 +cache[text1Id-1][text2Id]
            cache[text1Id][text2Id]= max(case1, case2, case3)
        }
    }
    //fmt.Println(cache)
    return cache[len(text1)-1][len(text2)-1]
    //return f(text1,text2,len(text1)-1,len(text2)-1,cache)
}


func f(text1 string, text2 string, text1Id int, text2Id int,cache [][]int) int {
    if text1Id < 0 || text2Id < 0 {
        return 0
    }
    if cache[text1Id][text2Id] != -1 {
        return cache[text1Id][text2Id]
    }
    case1, case2, case3 := 0,0,0
    if text1[text1Id] == text2[text2Id] {
        case1 = 1+f(text1,text2,text1Id-1,text2Id-1,cache) 
    } 
    case2 = 0 +f(text1,text2,text1Id, text2Id-1,cache)
    case3 = 0 +f(text1,text2,text1Id-1, text2Id,cache)
    cache[text1Id][text2Id]= max(case1, case2, case3)
    return max(case1, case2, case3)

}