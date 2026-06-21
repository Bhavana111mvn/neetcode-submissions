func minRemoveToMakeValid(s string) string {
	toskip := make(map[int]bool)
    stack := make([]int, 0)
    finalstr := make([]byte,0)
    for i:=0; i< len(s); i++ {
        if  s[i] == '(' {
            stack = append(stack,i)
        } else if s[i] == ')' {
            if len(stack) >0  {
                stack = stack[:len(stack)-1]
            } else {
                toskip[i] = true
            }
        }
    }
	for _, val:= range stack {
		toskip[val] = true
	}
    for i:=0; i< len(s); i++ {
        if !toskip[i] {
            finalstr = append(finalstr, s[i])
        }
    }
    return string(finalstr)
}
