func minOperations(s string) int {
    cnt := 0
    for i := 0; i < len(s); i++ {
        // For pattern starting with '0': even index should be '0', odd index should be '1'
        if i%2 == 0 {
            if s[i] == '1' {
                cnt++
            }
        } else {
            if s[i] == '0' {
                cnt++
            }
        }
    }
    // min(cnt, len(s)-cnt)
    if cnt < len(s)-cnt {
        return cnt
    }
    return len(s) - cnt
}