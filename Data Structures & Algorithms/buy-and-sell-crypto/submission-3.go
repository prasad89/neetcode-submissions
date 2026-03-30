func maxProfit(prices []int) int {
	maxProfit:=0
	for i:=0;i<len(prices)-1;i++{
		for j:=i+1;j<=len(prices)-1;j++{
			maxProfit=max(maxProfit,prices[j]-prices[i])
		}
	}
	return maxProfit
}
