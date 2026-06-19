func isValid(s string) bool {
    if len(s)%2 != 0 {
        return false
    }

	bracketMap := map[byte]byte{
        ')': '(',
        '}': '{',
        ']': '[',
    }

	var stack []byte

	for i := 0; i < len(s); i++ {
		if v, exists := bracketMap[s[i]]; exists{
			if len(stack) == 0 || stack[len(stack)-1] != v {
				return false
			}
			stack = stack[:len(stack)-1]
			continue
		} 
		stack = append(stack, s[i])
	}

	if len(stack) != 0 {
		return false
	}
	return true
}
