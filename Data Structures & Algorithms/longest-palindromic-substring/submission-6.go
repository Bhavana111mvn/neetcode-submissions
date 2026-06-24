func longestPalindrome(s string) string {
    res:=string(s[0])
	resLen := 0
	if len(s)==1 {
		return s
	}
	for start:=0; start<len(s);start++ {
		/// odd palindrome
		left, right:=start-1, start+1
		flag := false
		for left>=0 && right<len(s) && s[left]==s[right] {
			left--
			right++
			flag = true
		}
		left++
		right--
		if flag && resLen <= (right-left+1){
			resLen = right-left+1
			res = s[left:right+1]
			fmt.Println(start)
			fmt.Println(resLen)
		}
		/// even palindrome
		left, right = start, start+1
		flag = false
		for left>=0 && right<len(s) && s[left]==s[right] {
			left--
			right++
			flag = true
		}
        left++
		right--
		
		if flag && resLen <= (right-left+1){
			resLen = right-left+1
			res = s[left:right+1]
			fmt.Println(start)
			fmt.Println(resLen)
		}
	}
	return res
}

