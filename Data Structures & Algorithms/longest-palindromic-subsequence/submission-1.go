func longestPalindromeSubseq(s string) int {
    s1 := make([]byte, len(s))
	left :=0
	for i:=len(s)-1;i>=0;i-- {
        s1[left] = s[i]
		left++
	}

	return longestCommonSubsequence(string(s1),s)
}
func longestCommonSubsequence(text1 string, text2 string) int {
    if len(text1) == 0 || len(text2) == 0 {
        return 0
    }

    cache := make([][]int, len(text1)+1)
    for i := range cache {
        cache[i] = make([]int, len(text2)+1)
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

    for text1Id := 1; text1Id < len(text1); text1Id++ {
        for text2Id := 1; text2Id < len(text2); text2Id++ {
            case1, case2, case3 := 0, 0, 0

            if text1[text1Id] == text2[text2Id] {
                case1 = 1 + cache[text1Id-1][text2Id-1]
            }

            case2 = cache[text1Id][text2Id-1]
            case3 = cache[text1Id-1][text2Id]

            cache[text1Id][text2Id] = max(case1, max(case2, case3))
        }
    }

    return cache[len(text1)-1][len(text2)-1]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}