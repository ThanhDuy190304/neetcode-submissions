func minWindow(s string, t string) string {
	if (len(t) > len(s)){
		return ""
	}

    idx := func(c byte) int {
		if c >= 'a' {
			return int(c-'a') + 26  
		}
    	return int(c - 'A')         
	}

	var tCount, window [52]int
	inc := func(c byte) { window[idx(c)]++ }
	dec := func(c byte) { window[idx(c)]-- }

	satisfied := func() bool {
		for i := 0; i < 52; i++ {
			if window[i] < tCount[i] {
				return false
			}
		}
		return true
	}

	for i := 0; i < len(t); i++ {
    	tCount[idx(t[i])]++
    	inc(s[i])
	}

	if satisfied(){
		if(len(t) < len(s)){
			return t
		}
		return s
	}

	l, r := 0, len(t)
	var result [2]int 
	hasResult := false

	for (r < len(s)){
		inc(s[r])
		fmt.Println("A, r = ", r, "\n")

		for (l <= r && window[idx(s[l])] > tCount[idx(s[l])] ){
			dec(s[l])
			fmt.Println("B, l = ", l, "\n")
			l++
		}
		if satisfied() {
			if !hasResult || r-l < result[1]-result[0] {
				result[0] = l
				result[1] = r
			}
			hasResult = true
		}
		r++
	}

	if hasResult{
		fmt.Println("l = ", l, "r = ", r)
		return s[result[0]: result[1]+1]
	}
	return ""

}	
