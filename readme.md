# Regular Expression Engine using the Thompson Construction

>Author: Bastian Graebener\
>Student-ID: G00340600\
>Module: Graph Theory\
>Year: 3\
>Program: Bsc. in Software Development\
>Lecturer: Ian McLoughlin\
>Institute: Galway-Mayo Institute of Technology

## 1. Introduction
This program is a programming excersise for the module 'Graph Theory' in a 3rd year Software Development Course.

The task was to write a simple regular expression matching application that is based on Ken Thompsons' construction.

Thompsons' construction is used to convert a regular expression [[1](https://en.wikipedia.org/wiki/Regular_expression), 
    [2](https://www.regular-expressions.info)] to a 
Nondeterministic Finite Automaton(NFA) [[1](https://people.cs.clemson.edu/~goddard/texts/theoryOfComputation/3a.pdf), 
    [2](https://en.wikipedia.org/wiki/Nondeterministic_finite_automaton)].
This NFA is then used to match a String against the original regular expression.  

Minimum requirement for the application was to accept a regular expression with the special characters 
".", "|" and "*" and to match it against an input string. The application had to be written in the Go language
and no regular expression libraries were to be used. 



## 2. Elements of the Application
### 2.1 Nondeterministic Finite Automaton (NFA)
A NFA is a state machine which is defined by 5 tuples. A finite set of states, a finite set of
input strings, a transition function, one initial state and a set of accepting states.

In comparison to deterministic finite automatons in which every transition from state to state is unique,
in a NFA there can be multiple ways to transition to multiple states. Therefor it is possible to be in multiple
states at the same time. This is the property that makes it nondeterministic.

### 2.2 Regular Expressions
A regular expression is a series of characters, some of which may have a special meaning.
For example, the three characters ".", "|", and "∗" have the special meanings “concatenate”, “or”, and "Kleene
star" respectively. So, 0.1 means a 0 followed by a 1, 0|1 means a 0 or a 1,
and 1∗ means any number of 1’s. 

### 2.3 Thompsons' Construction 
Thompsons construction is an algorithm developed by Ken Thompson in 1968.

The algorithm splits a regular expression into its smallest sub-expression. For every sub-expression a NFA is created. 
All those NFA are then put together into a single NFA which can be used to match a string.

An in depth explanation of the algorithm can be found [here](https://en.wikipedia.org/wiki/Thompson%27s_construction) and 
[here](https://swtch.com/~rsc/regexp/regexp1.html).


## 3. The Structure of the Application

The application works in three steps.

1. Rewrite the regular expression from infix notation to postfix notation using the shunting yard algorithm.
1. Create NFAs for every sub fragment in the regular expression and assemble them to a single NFA (Thompsons' construction).  
1. Check if the regular expression matches an input string.


#### 3.1 Infix to Postfix
The Thompson construction works best with regular expression that are written in postfix notation rather than in infix
notation. 

>>Example for infix notation: a.b.c -> The operators are between the the operands.\
>>Postfix notation: ab.c. Here the operators come after the two operands. 

The advantage of the postfix notation is that brackets can be omitted if operators with lower precedence need to 
be applied before operators with higher precedence.
>> E.g infix: a.(b|d)?  -> postfix: abd|?.

The conversion from infix to postfix regular expressions is done in the _infixToPostfix_ function.
This function uses the [Shunting yard algorithm](https://en.wikipedia.org/wiki/Shunting-yard_algorithm).\
This algorithm uses a LIFO-Stack to cache operators and special characters. All non-special characters are appended directly to an 
output data structure. 

When a closing bracket is encountered all elements with a lower precedence than the top element
on the stack are added to the output. The opening brackets are written to the stack but not appended to the output later.

A regular expression already in postfix notation is simply returned unchanged.

#### 3.2 Regular Expression to NFA
The function _regexToNfa_ converts a regular expression in postfix notation into a Nondeterministic Finite Automaton.
The struct 'state' represents a single state in a state machine. Every state has two edges.
The struct 'nfa' represents a NFA with a state for input and one as accept state.

The function parses every symbol in the regular expression and creates a single NFA for each symbol.
The application accepts the following operators in order of precedence: '*', '+','?', '.', '|'.
Each operator results in a different NFA. 

The '*' **operator** matches zero or any number of a character. One edge of the accept state points back to the initial 
state creating a loop accepting any number of the character, the other points to the accept state possibly accepting 
zero occurrences. 

The **'+' operator** matches one or more of the same character. One outgoing edge of the NFA points back to previous NFA and
the other edge points to the accept state. The new NFA has the initial state of the previous NFA and the newly 
created accept state.

The **'?' operator** matches zero or one occurrence of a character. The initial state has two edges, one going straight
to the accept state, thus allowing for zero numbers of the character. The other edge is connected to the previous NFA.
The accept state of the previous NFA points straight to the new accept state. This prevents matching of multiple 
occurrences of the same character. 

The **'.' operator** concatenates two characters by connecting the accepting state of the first character with the input state of
the second state. The new NFA gets the initial state of the first characters' NFA and the accepting state of the second
characters' NFA. 

Teh **'|' operator** allows for matching one of two characters. The initial state of the new NFA points with both
edges to the initial states of the two possible NFAs. The new accept state connects both accept states from the two 
possible characters.

After the whole regular expression is parsed all individual NFAs are now connected like a Linked List.

#### 3.3 Matching a String
The function `matches` is used to check whether a string matches a regular expression. It accepts a regular expression 
string and an input string. 

The function first calls the _infixToPostfix_ function to convert the regular expression. It then calls the _regexToNfa_
function to retrieve the NFA used to match the string.

It maintains two lists of states. One list is used to keep track of the states that are currently visited. The other
list stores the states that can be reached from all current states. 

It iterates over the input string. For each character of the input string it check the states that are currently
visited. If the symbol of the state matches the current character in the string all possible next states are added to
the list of next states. This list then becomes the new current list and the list of next states is cleared.

After the whole input string has been processed, the list of current states is lopped over to check if one of the 
current states is an accept state. If this is the case, the function returns a true value otherwise a false. 


## 4. How to run the application
Clone the repository
```git 
git clone https://github.com/bGraebener/RegExpEngine.git
```
```
cd %GOPATH%/regexpengine
```
Run the application
```go
go run main.go
```
 
## 5. Examples
`regular expression: a.b?.(c|d*)`\
`input string: adddd`\
Output: _The input string matches the regular expression._

`regular expression: a.b?.(c|d*)`\
`input string: abcccc `\
Output: _The input string does not match the regular expression._

`regular expression: 1+.2.3`\
`input string: 111123`\
Output: _The input string matches the regular expression._

`regular expression: 1+.2.3`\
`input string: 23`\
Output: _The input string does not match the regular expression._

`regular expression: 1?.2*.6`\
`input string: 16`\
Output: _The input string matches the regular expression._

## 6. References
#### 6.1 NFA's
https://people.cs.clemson.edu/~goddard/texts/theoryOfComputation/3a.pdf\
https://en.wikipedia.org/wiki/Nondeterministic_finite_automaton

#### 6.2 Regular Expressions
https://en.wikipedia.org/wiki/Regular_expression\
https://www.regular-expressions.info

#### 6.3 Thompsons' Construction
https://swtch.com/~rsc/regexp/regexp1.html 

#### 6.4 Shunting Yard
https://en.wikipedia.org/wiki/Shunting-yard_algorithm