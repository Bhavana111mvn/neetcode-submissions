func validWordSquare(words []string) bool {
    for j:=0; j<len(words[0]); j++ {
		tmpstr := ""
		for i:=0; i<len(words);i++ {
			if j < len(words[i]) {
				tmpstr = tmpstr + string(words[i][j])
			}
	    } 
		if tmpstr != words[j] {
			return false
		}
	}
	return true
	
}
