package main

import (
	"fmt"
)

/**
RegExpEngine

Author: Bastian Graebener - G00340600

Simple regular expression engine in GoLang.
Project for Graph Theory Module in Year 3 of GMIT Bsc. Software Development.

main.go - Main file
 */

func main() {
	fmt.Print(infixToPostfix("a.b.c*"))
}

//converts an infix regular expression to a postfix one
//e.g.
//a.b -> ab.
//a.(bb)*.c ->abb.*.c.
func infixToPostfix(original string) string {
	var postFix []rune
	operators := map[rune]int{'*': 10, '.': 9, '|': 8}

	//make a stack
	var stack []rune

	//convert string to runes
	for _, r := range original {

		//check for rune type
		switch {

		//add open bracket to stack
		case r == '(':
			stack = append(stack, r)

			//add everything from the stack to postfix until you encounter an opening bracket
		case r == ')':
			var x rune
			for stack[len(stack)-1] != '(' {
				stack, x = pop(stack)
				postFix = append(postFix, x)
			}
			//pop the remaining bracket
			stack, _ = pop(stack)

			//while something is on the stack and the precedence is less than the top element of the stack
			//add it to postfix
		case operators[r] > 0:
			for len(stack) > 0 && operators[r] <= operators[stack[len(stack)-1]] {
				var x rune
				stack, x = pop(stack)
				postFix = append(postFix, x)
			}
			stack = append(stack, r)

		default:
			//add word rune to the postfix slice
			postFix = append(postFix, r)
		}
	}

	//add anything left on the stack to postfix
	for len(stack) > 0 {
		var x rune
		stack, x = pop(stack)
		postFix = append(postFix, x)
	}

	return string(postFix)
}

//returns the top element of a slice and removes it from the stack
func pop(stack []rune) ([]rune, rune) {
	//get the top element
	r := stack[len(stack)-1]

	//remove it from the stack
	stack = stack[:len(stack)-1]

	return stack, r
}
