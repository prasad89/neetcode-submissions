func maxProfit(prices []int) int {
	maxProfit:=0
	minBuy:=prices[0]
	for i:=1;i<=len(prices)-1;i++{
		maxProfit=max(maxProfit, prices[i]-minBuy)
		minBuy=min(minBuy,prices[i])
	}
	return maxProfit
}