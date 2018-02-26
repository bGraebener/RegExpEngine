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
	fmt.Print(infixToPostfix("a.b.c"))
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

	}

	//concatenate two character runes followed by an operator rune


	return string(postFix)
}

//returns the first element of a slice and removes it from the stack
func pop(stack []rune) (rune, []rune) {
	//get the top element
	r := stack[0]

	//remove it from the stack
	stack = stack[1:]

	return r, stack
}
