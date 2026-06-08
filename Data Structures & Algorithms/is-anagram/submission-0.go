func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := [26]int{}
	n := len(s)

	for i := 0; i < (n+1)/2; i++ {
		last := n - i - 1

		count[s[i]-'a']++
		count[t[i]-'a']--

		if last != i {
			count[s[last]-'a']++
			count[t[last]-'a']--
		}
	}

	for _, value := range count {
		if value != 0 {
			return false
		}
	}

	return true
}