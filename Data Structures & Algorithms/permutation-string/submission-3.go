func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }

    s1Count := make([]int, 26)
    s2Count := make([]int, 26)

    for i := 0; i < len(s1); i++ {
        s1Count[s1[i]-'a']++
    }

    l := 0
    for r := 0; r < len(s2); r++ {
        // Mở rộng window: thêm s2[r] vào
        s2Count[s2[r]-'a']++

        // Nếu s2[r] bị dư so với s1 → thu hẹp từ trái
        // cho đến khi count về bằng hoặc l vượt qua r
        for s2Count[s2[r]-'a'] > s1Count[s2[r]-'a'] {
            s2Count[s2[l]-'a']--
            l++
        }

        // Window hợp lệ và đúng kích thước → tìm thấy
        if r-l+1 == len(s1) {
            return true
        }
    }

    return false
}