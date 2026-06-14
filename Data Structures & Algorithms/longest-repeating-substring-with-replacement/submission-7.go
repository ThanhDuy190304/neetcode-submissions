func characterReplacement(s string, k int) int {
    counts := make([]int, 26)
    
    inc := func(c byte) { counts[c-'A']++ }
    dec := func(c byte) { counts[c-'A']-- }

    l, r := 0, 0
    maxFreq := 0
    result := 0

    for r < len(s) {
        inc(s[r])
        
        if counts[s[r]-'A'] > maxFreq {
            maxFreq = counts[s[r]-'A']
        }

        for (r - l + 1) - maxFreq > k {
            dec(s[l]) 
            l++      
        }

        if r - l + 1 > result {
            result = r - l + 1
        }

        r++
    }

    return result
}