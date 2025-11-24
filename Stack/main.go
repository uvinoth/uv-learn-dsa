package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// If open and close Parentheses is same Parentheses type like [] and (), not like this [) {]
func isValid(s string) bool {
	stack := []rune{} // Create empty stack

	for _, ch := range s { // Itrate the give string
		// fmt.Println(ch)
		switch ch {
		case '{':
			stack = append(stack, '}')
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		default:
			// fmt.Println(stack[len(stack)-1])
			if len(stack) == 0 || stack[len(stack)-1] != ch {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// Why do we do currNum = currNum*10 + int(ch-'0')?
// To convert a sequence of digit characters ('1', '2', '3'...) into the actual integer number (123).
// You read a number digit by digit, so you must "shift left" the previous digits by multiplying by 10 and then add the new digit.
// int(ch - '0')   => convert char to its numeric value
// | ch  | ch-'0' | currNum = currNum*10 + (ch-'0') |
// | --- | ------ | ------------------------------- |
// | '1' | 1      | 0*10 + 1 = **1**                |
// | '2' | 2      | 1*10 + 2 = **12**               |
// | '3' | 3      | 12*10 + 3 = **123**             |

func decodeString(s string) string {
	countStack := []int{}
	strStack := []string{}

	curNum := 0
	curStr := ""

	for _, ch := range s {
		if unicode.IsDigit(ch) {
			curNum = curNum*10 + int(ch-'0')
		} else if ch == '[' {
			strStack = append(strStack, curStr)
			countStack = append(countStack, curNum)
			curNum = 0
			curStr = ""
		} else if ch == ']' {
			repetItem := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]

			prevStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			temp := ""

			for i := 0; i < repetItem; i++ {
				temp += curStr
			}
			curStr = prevStr + temp
		} else {
			curStr += string(ch)
		}

	}

	return curStr
}

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, ch := range tokens {
		switch ch {
		case "+", "-", "*", "/":
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			var res int
			switch ch {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				res = a / b
			}

			stack = append(stack, res)
		default:
			num, _ := strconv.Atoi(ch)
			stack = append(stack, num)

		}
	}

	return stack[len(stack)-1]
}

func longestValidParentheses(s string) int {
	stack := []int{-1}
	maxLen := 0

	for i, ch := range s {
		if ch == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				length := i - stack[len(stack)-1]

				if length > maxLen {
					maxLen = length
				}
			}
		}

	}

	return maxLen
}

func calculate(s string) int {
	stack := []int{}
	result := 0
	num := 0
	sign := 1

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			num = num*10 + int(ch-'0') // Build multi digit number
		} else if ch == '+' {
			result += sign * num
			num = 0
			sign = 1
		} else if ch == '-' {
			result += sign * num
			num = 0
			sign = -1
		} else if ch == '(' {
			stack = append(stack, result)
			stack = append(stack, sign)
			result = 0
			sign = 1
		} else if ch == ')' {
			result += sign * num
			num = 0

			sign = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result = prev + sign*result
		}
	}

	return result + sign*num
}

func nextGreater(nums []int) []int {
	n := len(nums)
	stack := []int{}      // Maintain large element
	res := make([]int, n) // Return the next generater element

	for i := n - 1; i >= 0; i-- { // Itrate reverse order because avoid time complecity
		if len(stack) > 0 && stack[len(stack)-1] <= nums[i] { // Remove the smaler element in the stack
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 { // If no element in stack add -1
			res[i] = -1
		} else {
			res[i] = stack[len(stack)-1]
		}

		stack = append(stack, nums[i])
	}

	return res
}

func main() {
	// problem :: Valid Paranthesis - (Leecode  20)
	// fmt.Println(isValid("()"))
	// fmt.Println(isValid("()[]{}"))
	// fmt.Println(isValid("(]"))

	// problem :: Decode String - (Leecode 394)
	// fmt.Println(decodeString("3[a]2[bc]"))     // aaabcbc
	// fmt.Println(decodeString("3[a2[c]]"))      // accaccacc
	// fmt.Println(decodeString("2[abc]3[cd]ef")) // abcabccdcdcdef

	// problem :: Evaluate Reverse Polish Notation - (Leecode 150)
	// tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	// fmt.Println(evalRPN(tokens))

	// problem :: Longest Valid Parentheses - (Leecode 32)
	// s := ")()())"
	// fmt.Println(longestValidParentheses(s))

	// problem ::  Basic Calculators - (Leecode 224)
	// s := "(1+(4+5+2)-3)+(6+8)"
	// fmt.Println(calculate(s))

	// Monitonic Stack - Is store the elememt specific order
	// arr := []int{4, 5, 2, 25}
	// res := nextGreater(arr)
	// fmt.Println(res)

	// problem ::  Daily Temperatures - (Leecode 739)
	temp := []int{73, 74, 75, 71, 69, 72, 76, 73}
	result := dailyTemperatures(temp)
	fmt.Println(result)
}

func dailyTemperatures(temp []int) []int {
	n := len(temp)        // get the length
	stack := []int{}      // Define the stack
	res := make([]int, n) // Define the result

	for i := n - 1; i >= 0; i-- {

		for len(stack) > 0 && temp[stack[len(stack)-1]] <= temp[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			res[i] = 0
		} else {
			res[i] = stack[len(stack)-1] - i
		}

		stack = append(stack, i)
	}

	return res
}
