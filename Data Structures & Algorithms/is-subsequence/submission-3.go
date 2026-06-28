func isSubsequence(s string, t string) bool {
   left := 0
   if s == "" {
	 return true
   }
    i :=0
	for left < len(t) && i<len(s){
		if s[i] == t[left] {
			i++
		}
		left++
	}  
   
   return i==len(s)

}
