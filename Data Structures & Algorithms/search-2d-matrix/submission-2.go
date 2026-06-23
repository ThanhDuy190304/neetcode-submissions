func searchMatrix(matrix [][]int, target int) bool {
    n := len(matrix)
    m := len(matrix[0])
    l, r := 0, m*n-1
    for l <= r {
        mid := (l + r) / 2
        val := matrix[mid/m][mid%m]
        if val < target {
            l = mid + 1
        } else if val > target {
            r = mid - 1
        } else {
            return true
        }
    }
    return false
}