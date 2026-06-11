func maxProfit(prices []int) int {
	buyPrice := prices[0]
	result := 0

	for i := 1; i < len(prices); i++{
		rev := prices[i] - buyPrice
		if(rev > result){
			result = rev
			continue
		} 
		if(prices[i] < buyPrice){
			buyPrice = prices[i] 
			continue
		}
	}

	return result
	
}
