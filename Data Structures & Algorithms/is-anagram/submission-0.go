func isAnagram(s string, t string) bool {
	if len(s)!=len(t){
		return false
	}
	seenS:=make(map[rune]int)
	for _,c:=range s{
		seenS[c]++
	}
	for _,c:=range t{
		seenS[c]--
	}
	for _,c:=range s{
		if count,_:=seenS[c];count!=0{
			return false
		}
	}
	return true
}
