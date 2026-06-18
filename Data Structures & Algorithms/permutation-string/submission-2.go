func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }

    var s1Count, window [26]int

    // Khởi tạo window đầu tiên
    for i := 0; i < len(s1); i++ {
        s1Count[s1[i]-'a']++
        window[s2[i]-'a']++
    }

    if s1Count == window {
        return true
    }

    // Trượt window, mỗi bước add s2[r] và remove s2[r - len(s1)]
    for r := len(s1); r < len(s2); r++ {
        window[s2[r]-'a']++
        window[s2[r-len(s1)]-'a']--
        if s1Count == window {
            return true
        }
    }

    return false
}