func minWindow(s string, t string) string {
    if len(t) > len(s) {
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

    need := 0
    for i := 0; i < len(t); i++ {
        if tCount[idx(t[i])] == 0 {
            need++
        }
        tCount[idx(t[i])]++
    }

    have := 0
    l := 0
    var result [2]int
    hasResult := false

    for r := 0; r < len(s); r++ {
        inc(s[r])
        if tCount[idx(s[r])] > 0 && window[idx(s[r])] == tCount[idx(s[r])] {
            have++
        }

        for have == need {
            if !hasResult || r-l < result[1]-result[0] {
                result[0], result[1] = l, r
                hasResult = true
            }
            dec(s[l])
            if tCount[idx(s[l])] > 0 && window[idx(s[l])] < tCount[idx(s[l])] {
                have--
            }
            l++
        }
    }

    if !hasResult {
        return ""
    }
    return s[result[0] : result[1]+1]
}