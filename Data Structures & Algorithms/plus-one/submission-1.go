func plusOne(digits []int) []int {
    carry:=1
	for i:=len(digits)-1;i>=0;i--{
		sum:=digits[i]+carry
		if sum>9{
			digits[i]=0
		}else{
			digits[i]=sum
			return digits
		}
	}
	return append([]int{1},digits...)
}
