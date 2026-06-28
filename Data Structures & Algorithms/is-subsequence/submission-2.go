func isSubsequence(s string, t string) bool {
   left := 0
   flag := false
   if s == "" {
	 return true
   }
   for i:=0; i<len(s); i++ {
	   flag = false
	   for left < len(t) {
		  if s[i] == t[left] {
			flag = true
			left++
			break
		  }
		  left++
	    }  
   }
   return flag

}
