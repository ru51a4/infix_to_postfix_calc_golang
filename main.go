package main

import (
	"fmt"
	"strconv"

	"golang.org/x/exp/slices"
)

type calc struct{}

func (calc calc) lex(str string) []string {
	var res []string
	var t string
	for _, ch := range str {
		if ch == ' ' {
			continue
		}
		if ch == '/' || ch == '*' || ch == '-' || ch == '+' {
			if len(t) > 0 {
				res = append(res, t)
			}
			res = append(res, string(ch))
			t = ""
		} else if ch == '(' || ch == ')' {
			if len(t) > 0 {
				res = append(res, t)
			}
			res = append(res, string(ch))
			t = ""
		} else {
			t = t + string(ch)
		}
	}
	if len(t) > 0 {
		res = append(res, t)
	}
	return res
}
func (calc calc) convert(arr []string) []string {
	var stack []string
	var stackOp []string
	op1 := []string{"+", "-"}
	op2 := []string{"*", "/"}
	for len(arr) > 0 {
		c := arr[0]
		arr = arr[1:]
		if len(stackOp) == 0 && (slices.Contains(op1, c) || slices.Contains(op2, c)) {
			stackOp = append(stackOp, c)
			continue
		}
		var lastOp string
		if len(stackOp) > 0 {
			lastOp = stackOp[len(stackOp)-1]
		}
		if c == "(" {
			stackOp = append(stackOp, c)
		} else if c == ")" {
			q := stackOp[len(stackOp)-1]
			stackOp = stackOp[:len(stackOp)-1]
			for q != "(" {
				stack = append(stack, q)
				q = stackOp[len(stackOp)-1]
				stackOp = stackOp[:len(stackOp)-1]
			}
		} else if slices.Contains(op1, c) {
			if slices.Contains(op1, lastOp) {
				stack = append(stack, stackOp[len(stackOp)-1])
				stackOp = stackOp[:len(stackOp)-1]
				stackOp = append(stackOp, c)
			}
			if slices.Contains(op2, lastOp) {
				stack = append(stack, stackOp[len(stackOp)-1])
				stackOp = stackOp[:len(stackOp)-1]
				stackOp = append(stackOp, c)
			}
			if lastOp == "(" {
				stackOp = append(stackOp, c)
			}
		} else if slices.Contains(op2, c) {
			if slices.Contains(op1, lastOp) {
				stackOp = append(stackOp, c)
			}
			if slices.Contains(op2, lastOp) {
				stack = append(stack, stackOp[len(stackOp)-1])
				stackOp = stackOp[:len(stackOp)-1]
				stackOp = append(stackOp, c)
			}
			if lastOp == "(" {
				stackOp = append(stackOp, c)
			}
		} else {
			stack = append(stack, c)
		}
	}

	for len(stackOp) > 0 {
		stack = append(stack, stackOp[len(stackOp)-1])
		stackOp = stackOp[:len(stackOp)-1]
	}
	return stack
}
func (calc calc) eval(arr []string) string {
	var stack []string
	for len(arr) > 0 {
		c := arr[0]
		arr = arr[1:]
		if c == "+" {
			n2, _ := strconv.Atoi(stack[len(stack)-1])
			n1, _ := strconv.Atoi(stack[len(stack)-2])
			stack = stack[:len(stack)-2]
			stack = append(stack, strconv.Itoa(n1+n2))
		} else if c == "-" {

			n2, _ := strconv.Atoi(stack[len(stack)-1])
			n1, _ := strconv.Atoi(stack[len(stack)-2])
			stack = stack[:len(stack)-2]
			stack = append(stack, strconv.Itoa(n1-n2))
		} else if c == "/" {
			n2, _ := strconv.Atoi(stack[len(stack)-1])
			n1, _ := strconv.Atoi(stack[len(stack)-2])
			stack = stack[:len(stack)-2]
			stack = append(stack, strconv.Itoa(n1/n2))
		} else if c == "*" {
			n2, _ := strconv.Atoi(stack[len(stack)-1])
			n1, _ := strconv.Atoi(stack[len(stack)-2])
			stack = stack[:len(stack)-2]
			stack = append(stack, strconv.Itoa(n1*n2))
		} else {
			stack = append(stack, c)
		}
	}
	return stack[0]
}

func main() {
	c := calc{}
	var kek string = c.eval(c.convert(c.lex("4+4+2*2")))
	fmt.Println(kek)
}
