func evalRPN(tokens []string) int {
	operators := map[string]struct{}{
    	"+": {},
    	"-": {},
    	"*": {},
    	"/": {},
	}
	calculate := func(num1, num2 int, operator string) int{
		if operator == "+"{
			return num1 + num2
		}else if operator == "-"{
			return num1 - num2
		}else if operator == "*"{
			return num1 * num2
		}else{
			return num1 / num2
		}
	}
	var stack []int

	for _, t := range tokens{
		if _, exists := operators[t]; exists {
			num1 := stack[len(stack) - 2]
			num2 := stack[len(stack) - 1]
			result := calculate(num1, num2, t)
			fmt.Println(num1, " ", num2, " ", result )
			stack = append(stack[:len(stack) - 2], result)
			continue
		}
		val, _ := strconv.Atoi(t)
		stack = append(stack, val)
	}

	return stack[len(stack) - 1]


}
