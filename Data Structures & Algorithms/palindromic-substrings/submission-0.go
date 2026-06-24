func countSubstrings(s string) int {
    cnt := 0
    for start:=0; start<len(s);start++ {
        cnt+=1
        left, right := start-1,start+1
        for left>=0 && right<len(s) && s[left]==s[right] {
            cnt++
            left--
            right++
        }

        left, right = start, start+1
        for left>=0 && right<len(s) && s[left]==s[right] {
            cnt++
            left--
            right++
        }
    }
    return cnt
}
