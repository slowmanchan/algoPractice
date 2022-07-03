package main

func main() {
}

// array stack
func isValid(str string) bool {
	// init a stack
	stack := []string{}

	// memoize the close open pairs for fast lookup
	memo := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	// loop thru each character in str
	for _, c := range str {
		// check the map to see if its a closing bracket
		if _, ok := memo[string(c)]; ok {
			// get the top element of the stack
			// all brackets need to be resolved in order
			// you cant have {[}]
			topElement := ""
			if len(stack) > 0 {
				// no pop function in golang (sadness)
				topElement = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}

			// check the value of the top element
			// against the needed pair in the mapping
			// ) needs a (. if it doesnt match, we've got
			// an invalid brackets
			if memo[string(c)] != topElement {
				return false
			}
		} else {
			// if its not in the mapping, in this case its an
			// open bracket because the mapping only has closed brackets as
			// keys. In the case we add it to our stack to process later.
			// this keeps the ordering when we have to check valid closed parts
			stack = append(stack, string(c))
		}
	}

	// if the stack still has elements, then we know its invalid
	// because a pair wasnt found for it
	return len(stack) <= 0
}
