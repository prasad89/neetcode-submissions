func isPalindrome(s string) bool {
	i,j:=0,len(s)-1
	for i<=j{
		fmt.Println(string(s[i]),string(s[j]))
		if !isAlphanumeric(rune(s[i])) {
			i++
			continue
		}
		if !isAlphanumeric(rune(s[j])){
			j--
			continue
		}
		if strings.ToLower(string(s[i]))!=strings.ToLower(string(s[j])){
			return false
		}
		i++
		j--
	}
	return true
}

func isAlphanumeric(c rune)bool{
	return unicode.IsLetter(c)||unicode.IsDigit(c)
}


// nolemonnomelon
// nolemonnomelon