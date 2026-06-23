func minEatingSpeed(piles []int, h int) int {
	
	maxPile := 0
	total := 0
	for _, pile := range(piles){
		if pile > maxPile{
			maxPile = pile
		}
		total += pile
	}

	l, r := (total + h - 1) / h , maxPile
	fmt.Println("l:", l, "r:", r, "total:", total, "maxPile:", maxPile)
	for (l < r){
		mid := (l + r) / 2
		fmt.Println("--- mid:", mid)
		tmp := 0
		for _, pile := range(piles){
			tmp += (pile + mid - 1) / mid
			if(tmp > h){
				break
			}
		}
		fmt.Println("tmp:", tmp, "h:", h)
		if tmp > h {
			l = mid + 1
		}else{
			r = mid 
		}
	}
	return l
}
