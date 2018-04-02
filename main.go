/**
RegExpEngine

Author: Bastian Graebener - G00340600

Simple regular expression engine in GoLang.
Project for Graph Theory Module in Year 3 of GMIT Bsc. Software Development.

main.go - Main file
*/

package main

import "fmt"

func main() {
	//fmt.Println(infixToPostfix("a.b|c*"))
	//
	//nfa := regexToNfa(infixToPostfix("a.b|c*"))
	//fmt.Println(nfa)

	fmt.Println(matches("a.b|c*", "cccc"))

}

//converts an infix regular expression to a postfix one
//e.g. a.b -> ab. or a.(bb)*.c ->abb.*.c.
func infixToPostfix(original string) string {
	//tmp variable
	var x rune
	//slice to collect output
	var postFix []rune

	//map of special characters and their weight
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
		stack, x = pop(stack)
		postFix = append(postFix, x)
	}

	return string(postFix)
}

//struct representing states and edges
type state struct {
	symbol     rune
	firstEdge  *edge
	secondEdge *edge
}

// alias for clarity
type edge = state

//struct representing a non-deterministic finite automaton
type nfa struct {
	initial *state
	accept  *state
}

//converts a regular expression in postfix notation into a nfa
func regexToNfa(postfix string) *nfa {
	var nfaStack []*nfa

	for _, r := range postfix {
		switch r {

		//concatenate two fragments
		case '.':
			//	pop two elements of the stack of regex fragments
			twoFragments := nfaStack[len(nfaStack)-2:]
			nfaStack = nfaStack[:len(nfaStack)-2]

			//connect accept state of firstEdge fragment to initial state of secondEdge fragment
			twoFragments[0].accept.firstEdge = twoFragments[1].initial

			//create a new state with the initial state of the firstEdge fragment and the accept state of the secondEdge fragment
			newNfa := &nfa{initial: twoFragments[0].initial, accept: twoFragments[1].accept}

			//and push it back to the stack
			nfaStack = append(nfaStack, newNfa)

			//	accept either fragments
		case '|':
			//	pop two elements of the stack of regex fragments
			twoFragments := nfaStack[len(nfaStack)-2:]
			nfaStack = nfaStack[:len(nfaStack)-2]

			//	create a new initial and a new accept state
			var initial state
			var accept state

			// let all initial states from both elements point to the new initial state
			initial.firstEdge = twoFragments[0].initial
			initial.secondEdge = twoFragments[1].initial

			//let all accept states from both elements point to the new accept state
			twoFragments[0].accept.firstEdge = &accept
			twoFragments[1].accept.firstEdge = &accept

			//push the new fragment back onto the stack
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})

			//	accept any number of on element
		case '*':
			//	pop one fragment of the stack
			frag := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			//create a new initial and accept state
			var initial state
			var accept state

			//let the new initial state point to the old initial and new accept state
			initial.firstEdge = frag.initial
			initial.secondEdge = &accept

			//let the old accept state point to the old initial state and the new accept state
			frag.accept.firstEdge = frag.initial
			frag.accept.secondEdge = &accept

			//push a new fragment containing the new states onto the stack
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})

			//	non-special characters
		default:
			//create two states
			var accept state
			//store the character
			initial := state{symbol: r, firstEdge: &accept}

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		}
	}
	return nfaStack[0]
}

//checks whether an infix regular expression matches an input
func matches(regex string, original string) bool {
	matches := false

	//convert infix regular expression to postfix expression
	regex = infixToPostfix(regex)

	// compile a regular expression to a NFA
	regexNfa := regexToNfa(regex)

	//keep track of all states that the algorithm is currently in
	var current []*state
	//keep track of all states that can be reached next
	var next []*state

	current = addState(current[:], regexNfa.initial, regexNfa.accept)

	//loop the string
	for _, r := range original {
		//check all states currently in and what next possible states to reach
		for _, c := range current {

			// if the current states' symbol matches the current character in the input string put it into the list
			// of next states
			if c.symbol == r {
				next = addState(next[:], c, regexNfa.accept)
			}
		}
		//swap the current states against the old next states and clear the next state slice
		current, next = next, []*state{}
	}

	//after the whole original string has been gone over, check if the nfa is in an accept state
	for _, c := range current {
		if c == regexNfa.accept {
			matches = true
			break
		}
	}

	return matches
}

// helper function to add states to a list of states
// also adds all states that can be reached from the newly added states
func addState(states []*state, first *state, accept *state) []*state {
	//add the immediate states
	states = append(states, first, accept)

	// 0 symbol means e-edges coming to the state
	// all the states the e-edges point to have to be added as well
	if first != accept && first.symbol == 0 {
		states = addState(states, first.firstEdge, accept)
		if first.secondEdge != nil {
			states = addState(states, first.secondEdge, accept)
		}
	}
	return states
}

//returns the top element of a slice and removes it from the stack
func pop(stack []rune) ([]rune, rune) {
	//get the top element
	r := stack[len(stack)-1]

	//remove it from the stack
	stack = stack[:len(stack)-1]

	return stack, r
}
