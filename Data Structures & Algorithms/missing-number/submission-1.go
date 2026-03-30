func missingNumber(nums []int) int {
	actualSum:=0
	for _,num:=range nums{
		actualSum+=num
	}
	expectedSum:=len(nums)*(len(nums)+1)/2
	return expectedSum-actualSum
}
