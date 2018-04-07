# Simple Regular Expression Engine

>Author: Bastian Graebener 
>
>Student-ID: G00340600
>
>Module: Graph Theory
>
>Year: 3
>
>Program: Bsc. in Software Development
>
>Lecturer: Ian McLoughlin
>
>Institute: Galway-Mayo Institute of Technology

## 1. Introduction
This program is a programming excersise for the module 'Graph Theory' in a 3rd year Software Development Course.

The task was to write a simple regular expression matching application that is based on Ken Thompsons' algorithm. An 
explanation of the algorithm can be found [here](https://en.wikipedia.org/wiki/Thompson%27s_construction) and 
[here](https://swtch.com/~rsc/regexp/regexp1.html).

Thompsons' algorithm is used to convert a regular expression [[1](https://en.wikipedia.org/wiki/Regular_expression), 
    [2](https://www.regular-expressions.info)] to a 
Nondeterministic Finite Automaton(NFA) [[1](https://people.cs.clemson.edu/~goddard/texts/theoryOfComputation/3a.pdf), 
    [2](https://en.wikipedia.org/wiki/Nondeterministic_finite_automaton)].
This NFA is then used to match a String against the original regular expression.  

## 2. 

### 2.1 Nondeterministic Finite Automaton


### 2.2 Regular Expressions
A regular expression is a series of characters, some of which may have a special meaning.
For example, the three characters ".", "|", and "∗" have the special meanings “concatenate”, “or”, and "Kleene
star" respectively. So, 0.1 means a 0 followed by a 1, 0|1 means a 0 or a 1,
and 1∗ means any number of 1’s. These special characters must be used in
your submission.

### 2.3 Thompsons' Algorithm 



## 3. The Application

1. Convert from infix to postfix (shunting yard)
1. Create nfas for every sub fragment in regular expression (Thompsons' Algorithm)  
1. match the regular expression against an input string


## 4. How to run the application




## 5. References
https://swtch.com/~rsc/regexp/regexp1.html 