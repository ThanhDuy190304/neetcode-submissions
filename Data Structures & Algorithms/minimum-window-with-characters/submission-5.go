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
    need := 0
    for i := 0; i < len(t); i++ {
        if tCount[idx(t[i])] == 0 {
            need++
        }
        tCount[idx(t[i])]++
    }

    have := 0
    l := 0
    resL, resR := 0, -1

    for r := 0; r < len(s); r++ {
        i := idx(s[r])
        window[i]++
        if tCount[i] > 0 && window[i] == tCount[i] {
            have++
        }

        for have == need {
            if resR == -1 || r-l < resR-resL {
                resL, resR = l, r
            }
            j := idx(s[l])
            window[j]--
            if tCount[j] > 0 && window[j] < tCount[j] {
                have--
            }
            l++
        }
    }

    if resR == -1 {
        return ""
    }
    return s[resL : resR+1]
}