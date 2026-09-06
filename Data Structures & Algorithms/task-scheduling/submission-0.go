func leastInterval(tasks []byte, n int) int {
    var count [26]int
    for _, t := range tasks {
        count[t-'A']++
    }
    sort.Sort(sort.Reverse(sort.IntSlice(count[:])))

    maxFreq := count[0]
    idle := (maxFreq - 1) * n
    for i := 1; i < 26 && count[i] > 0; i++ {
        idle -= min(count[i], maxFreq-1)
    }
    if idle < 0 {
        idle = 0
    }
    return len(tasks) + idle
}