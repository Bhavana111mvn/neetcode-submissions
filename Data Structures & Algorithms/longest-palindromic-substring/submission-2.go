func longestPalindrome(s string) string {
    res:=string(s[0])
	resLen := 0
	if len(s)==1 {
		return s
	}
	for i:=0; i<len(s);i++ {
		for j:=i+1;j<len(s);j++ {
			if chckisPalindrome(i,j,s) && resLen < (j-i+1){
				resLen = j-i+1
				res = s[i:j+1]
			}
		}
	}
	return res
}

func chckisPalindrome(l ,r int, s string)bool {
	for l <=r {
		if s[l]!=s[r] {
			return false
		}
		l++
		r--
	}
	return true
}