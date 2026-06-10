func isPalindrome(s string) bool {
	isValidChar := func(c byte) bool {
		if ('a' <= c && c <= 'z') || 
			('A' <= c && c <= 'Z') ||
			('0' <= c && c <= '9') {
			return true
		}
		return false
	}

	toLower := func(c byte) byte {
		if 'A' <= c && c <= 'Z' {
			return c + ('a' - 'A')
		}
		return c
	}

	left := 0
	right := len(s) - 1
	for {
		if (left >= right){
			break
		}
		
		if !isValidChar(s[left]){
			left ++
			continue 
		}
		
		if !isValidChar(s[right]){
			right--;
			continue;
		}

		if toLower(s[left]) != toLower(s[right]){
			return false
		}

		left ++
		right --
		
	}
	return true
}
