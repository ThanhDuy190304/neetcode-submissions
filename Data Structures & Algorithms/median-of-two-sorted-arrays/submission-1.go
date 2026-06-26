func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n :=	len(nums1), len(nums2)
	if (m > n){
		return findMedianSortedArrays(nums2, nums1)
	}

	l, r := 0, m 

	for l <= r {
		i := (l + r) / 2
		j := (m + n + 1) / 2 - i

		maxL1  := math.MinInt64
    	minR1 := math.MaxInt64
        maxL2  := math.MinInt64
        minR2 := math.MaxInt64
		
		if i > 0 { maxL1 = nums1[i-1] }
        if i < m { minR1 = nums1[i] }
        if j > 0 { maxL2 = nums2[j-1] }
        if j < n { minR2 = nums2[j] }

		if maxL1 <= minR2 && maxL2 <= minR1 {
            if (m+n)%2 == 1 {
                return float64(max(maxL1, maxL2))
            }
            return float64(max(maxL1, maxL2)+min(minR1, minR2)) / 2.0
        } else if maxL1 > minR2 {
            r = i - 1  
        } else {
            l = i + 1  
        }
	}
	return 0.0

}
