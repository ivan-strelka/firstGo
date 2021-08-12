package rpn_calc

func rpn(args ...interface{}) int {
	return rpn1(args)

}

func isInt(x interface{}) bool {
	_, ok := x.(int)
	return ok

}

// 3 2 7 + * == 27
// 3 [3]
// 2 [3 2]
// 7 [3 2 7]
// + [3 9]
// * [27]
func rpn1(args []interface{}) int {
	var stack []int
	for _, arg := range args {
		if isInt(arg) {
			stack = append(stack, arg.(int))
		}
		if arg == "-" {
			result := stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[0 : len(stack)-2]
			stack = append(stack, result)
		}
		if arg == "*" {
			result := stack[len(stack)-2] * stack[len(stack)-1]
			stack = stack[0 : len(stack)-2]
			stack = append(stack, result)
		}

		if arg == "+" {
			result := stack[len(stack)-2] + stack[len(stack)-1]
			stack = stack[0 : len(stack)-2]
			stack = append(stack, result)
		}
		if arg == "/" {
			result := stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[0 : len(stack)-2]
			stack = append(stack, result)
		}

	}
	return stack[len(stack)-1]

}
