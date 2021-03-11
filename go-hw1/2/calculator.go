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

type Stack struct {
	stack []interface{}
}

func (st *Stack) Push(value interface{}) {
	st.stack = append(st.stack, value)
}

func (st *Stack) Pop() interface{} {
	value := st.stack[len(st.stack)-1]
	st.stack = st.stack[:len(st.stack)-1]
	return value
}

func count(expression string) int32 {
	var number int32
	var flag bool
	var stack Stack

	for _, element := range expression {
		if unicode.IsDigit(element) {
			number *= 10
			number += element - '0'
			flag = true
		} else {
			if element != ' ' {
				num2 := stack.Pop().(int32)
				num1 := stack.Pop().(int32)

				stack.Push(choseOperation(element, num1, num2))
				flag = false
			} else if element == ' ' && flag {
				stack.Push(number)
				number = 0
			}
		}
	}
	return stack.stack[len(stack.stack)-1].(int32)
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
	var stack Stack

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
			if len(stack.stack) == 0 || priorities[stack.stack[len(stack.stack)-1].(string)] < priorities[strElem] {
				stack.Push(strElem)
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
				for priorities[stack.stack[len(stack.stack)-1].(string)] > priorities[strElem] {
					resultExp += stack.stack[len(stack.stack)-1].(string) + " "
					stack.Pop()
				}
				stack.Push(strElem)
			}
			flag = false
		} else if strElem == "(" {
			if endOfNumber(elem) && flag {
				resultExp += strconv.Itoa(int(number)) + " "
				number = 0
				flag = false
			}
			stack.Push(strElem)
			flag = false
		} else if strElem == ")" {
			if endOfNumber(elem) && flag {
				resultExp += strconv.Itoa(int(number)) + " "
				number = 0
				flag = false
			}
			for stack.stack[len(stack.stack)-1].(string) != "(" {
				resultExp += stack.stack[len(stack.stack)-1].(string) + " "
				stack.Pop()
			}
			stack.Pop()
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

	for len(stack.stack) != 0 {
		resultExp += stack.stack[len(stack.stack)-1].(string) + " "
		stack.Pop()
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
