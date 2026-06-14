func characterReplacement(s string, k int) int {
	result := 0
	for i := 0; i < len(s); i++ {
		cnt1 := 1
		cnt2 := 0
		j := i + 1

		for(j < len(s) && (cnt2 < k || (s[j] == s[i]))){
			if (s[j] == s[i]){
				cnt1 ++
				j++
				continue
			}
			cnt1 ++
			cnt2 ++
			j++
		}

		for(cnt2 < k && cnt1 < len(s)){
			cnt1++
			cnt2++
		}
	
		if (cnt1 > result){
			result = cnt1
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		cnt1 := 1
		cnt2 := 0
		j := i - 1

		for(j >= 0 && (cnt2 < k || (s[j] == s[i]))){
			if (s[j] == s[i]){
				cnt1 ++
				j--
				continue
			}
			cnt1 ++
			cnt2 ++
			j--
		}
		for(cnt2 < k && cnt1 < len(s)){
			cnt1++
			cnt2++
		}
		if (cnt1 > result){
			result = cnt1
		}
	}

	return result
}
