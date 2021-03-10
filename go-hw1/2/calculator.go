package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func choseOperation(operation, first, second int32) int32 {
	switch {
	case string(operation) == "-":
		return first - second
	case string(operation) == "+":
		return first + second
	case string(operation) == "/":
		return first / second
	}
	return first * second
}

func count(expression string) int32 {
	var number int32
	var flag bool
	var stack []int32

	for _, element := range expression {
		if unicode.IsDigit(element) {
			number *= 10
			number += element - '0'
			flag = true
		} else {
			if element != ' ' {
				num2 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				num1 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				stack = append(stack, choseOperation(element, num1, num2))
				flag = false
			} else if element == ' ' && flag {
				stack = append(stack, number)
				number = 0
			}
		}
	}
	return stack[len(stack)-1]
}

var (
	priorities = map[string]int{
		"*": 3,
		"/": 3,
		"+": 2,
		"-": 2,
		"(": 1,
	}
)

func endOfNumber(elem int32) bool {
	if elem == ' ' || elem == ')' || elem == '+' || elem == '-' ||
		elem == '*' || elem == '/' || elem == '(' {
		return true
	}
	return false
}

func formExpression(normalExpression string) string {
	var number int32
	var flag bool
	var resultExp string
	var stack []string

	for _, elem := range normalExpression {
		strElem := string(elem)
		if unicode.IsDigit(elem) {
			number *= 10
			number += elem - '0'
			flag = true
		} else if strElem == "-" || strElem == "/" ||
			strElem == "+" || strElem == "*" {
			if endOfNumber(elem) && flag {
				resultExp += strconv.Itoa(int(number)) + " "
				number = 0
				flag = false
			}
			if len(stack) == 0 || priorities[stack[len(stack)-1]] < priorities[strElem] {
				stack = append(stack, strElem)
				if (elem == ' ' || elem == ')') && flag {
					resultExp += strconv.Itoa(int(number)) + " "
					number = 0
					flag = false
				}
			} else {
				if endOfNumber(elem) && flag {
					resultExp += strconv.Itoa(int(number)) + " "
					number = 0
					flag = false
				}
				for priorities[stack[len(stack)-1]] > priorities[strElem] {
					resultExp += stack[len(stack)-1] + " "
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, strElem)
			}
			flag = false
		} else if strElem == "(" {
			if endOfNumber(elem) && flag {
				resultExp += strconv.Itoa(int(number)) + " "
				number = 0
				flag = false
			}
			stack = append(stack, strElem)
			flag = false
		} else if strElem == ")" {
			if endOfNumber(elem) && flag {
				resultExp += strconv.Itoa(int(number)) + " "
				number = 0
				flag = false
			}
			for stack[len(stack)-1] != "(" {
				resultExp += stack[len(stack)-1] + " "
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			flag = false
		} else if endOfNumber(elem) && flag {
			resultExp += strconv.Itoa(int(number)) + " "
			number = 0
			flag = false
		}
	}
	if number != 0 {
		resultExp += strconv.Itoa(int(number)) + " "
	}

	for len(stack) != 0 {
		resultExp += stack[len(stack)-1] + " "
		stack = stack[:len(stack)-1]
	}
	return resultExp
}

func readExpression() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return line
}

func main() {
	line := readExpression()
	fmt.Println(line)
	parsedExp := formExpression(line)
	fmt.Println(parsedExp)
	result := count(parsedExp)
	fmt.Println(result)
}
