func stringShift(s string, shift [][]int) string {
   runstr := []rune(s)
   for _, sh := range shift {
	    if sh[0] == 0 {
			/// left shift
			runstr = shiftele(runstr, sh[1], true, false)
		} else {
			//// right shift
			runstr = shiftele(runstr, sh[1], false, true)

		}
   }
   return string(runstr)
}

func shiftele(runstr []rune, cnt int, left bool, right bool)[]rune {
	if left {
		runstr = append(runstr[cnt:], runstr[0:cnt]...)
	} else {
		tmp := len(runstr)-cnt
		runstr = append(runstr[tmp:], runstr[0:tmp]...)
	}
	return runstr
}
