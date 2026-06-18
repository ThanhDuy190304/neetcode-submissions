func checkInclusion(s1 string, s2 string) bool {
	if (len(s1) > len(s2)) {
		return false;
	}
	s1Counts := make([]int, 26)
	s2SlideCounts := make([]int, 26)
	s2Inc := func(c byte) { s2SlideCounts[c-'a']++ }
	s2Dec := func(c byte) { s2SlideCounts[c - 'a']-- }

	for _, s := range s1 {
		s1Counts[s - 'a'] ++
	}
	l, r := -1, 0;
	for r < len(s2)  {
		fmt.Print("r = ", r)
		s2Inc(s2[r])
		if(s2SlideCounts[s2[r] - 'a'] > s1Counts[s2[r] - 'a']){
			fmt.Print("+")
			for (l <= r && s2SlideCounts[s2[r] - 'a'] > s1Counts[s2[r] - 'a']){
				if (l < 0){
					l++
					continue
				}
				s2Dec(s2[l])
				fmt.Print("l = ", l)
				l++
				if (l == len(s2)){
					break
				}
			}
			r++
			continue
		}
		if(r - l + 1 == len(s1) && l >= 0) || (r + 1 == len(s1) && l < 0) {
			
			return true
		}
		r++
	}

	return false

}
