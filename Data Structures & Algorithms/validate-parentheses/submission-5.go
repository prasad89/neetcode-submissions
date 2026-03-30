func isValid(s string) bool {
	if len(s)==1{
		return false
	}
    stack:=[]rune{}
	// LIFO
	for _,c:=range s{
		fmt.Println(stack,c)
		if c=='(' || c == '{' || c == '['{
			stack=append(stack,c)
			continue
		}
		if c == ')' && len(stack)!=0 && stack[len(stack)-1]=='('{
			stack=stack[:len(stack)-1]
		}else if c == '}' && len(stack)!=0 && stack[len(stack)-1]=='{' {
			stack=stack[:len(stack)-1]
		}else if c == ']' && len(stack)!=0 && stack[len(stack)-1]=='[' {
			stack=stack[:len(stack)-1]
		}else{
			return false
		}
	}
	return len(stack)==0
}
