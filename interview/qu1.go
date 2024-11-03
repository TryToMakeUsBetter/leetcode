package interview

/*
**\
({})
**
*/
func IfClose(input []byte) (res bool) {
	stack := make([]byte, len(input))
	top := -1
	for i := 0; i < len(input); i++ {
		if input[i] == '(' || input[i] == '{' || input[i] == '[' {
			top++
			stack[top] = input[i]
		}
		if input[i] == ')' {
			if stack[top] != '(' {
				return false
			} else {
				top--
			}
		}
		if input[i] == ']' {
			if stack[top] != '[' {
				return false
			} else {
				top--
			}
		}
		if input[i] == '}' {
			if stack[top] != '{' {
				return false
			} else {
				top--
			}
		}
	}
	return true
}
