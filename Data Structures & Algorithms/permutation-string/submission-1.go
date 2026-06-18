func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }

    s1Count := make([]int, 26)
    s2Count := make([]int, 26)
    inc := func(c byte) { s2Count[c-'a']++ }
    dec := func(c byte) { s2Count[c-'a']-- }

    for i := 0; i < len(s1); i++ {
        s1Count[s1[i]-'a']++
    }

    l, r := -1, 0
    for r < len(s2) {
        inc(s2[r])

        for l <= r && s2Count[s2[r]-'a'] > s1Count[s2[r]-'a'] {
            if l < 0 {
                l++
                continue
            }
            dec(s2[l])
            l++
        }

        if (r-l+1 == len(s1) && l >= 0) || (r+1 == len(s1) && l < 0) {
            return true
        }

        r++
    }

    return false
}