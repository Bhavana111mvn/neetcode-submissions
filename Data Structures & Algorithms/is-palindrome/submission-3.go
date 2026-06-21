func isPalindrome(s string) bool {
	bytarr := []rune(s)
	n:= len(bytarr)
    left, right := 0, n-1
	for left < right {
		for left < right && !isalphanumeric(bytarr[left]) {
			left++
		} 
		for left < right && !isalphanumeric(bytarr[right]) {
			right--
		} 
		if left >= right {
			break
	    }
		l, r := bytarr[left], bytarr[right]
	// convert to lowercase
		if l >= 'A' && l <= 'Z' {
			l += 32
	}
		if r >= 'A' && r <= 'Z' {
			r += 32
	}
		
		if l!= r {
			return false
	}
		left++
		right--
		
	}
	return true

}

func isalphanumeric(ch rune) bool{
	if (ch >= '0' && ch <= '9') || (65 <= ch && ch <= 90) || (97 <= ch && ch<=122) {
        return true
	}
	return false
}
