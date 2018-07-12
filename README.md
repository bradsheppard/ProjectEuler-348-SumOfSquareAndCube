# Sum of Square and Cube

This was built as a solution to problem 348 (Sum of a Square and a Cube) of Project Euler.

## Problem Description



Many numbers can be expressed as the sum of a square and a cube. Some of them in more than one way.

Consider the palindromic numbers that can be expressed as the sum of a square and a cube, both greater than 1, in exactly 4 different ways.
For example, 5229225 is a palindromic number and it can be expressed in exactly 4 different ways:

22852 + 203 \
22232 + 663 \
18102 + 1253 \
11972 + 1563 

Find the sum of the five smallest such palindromic numbers.


## Usage

The solution I implemented is written in GoLang. To build the program, execute the command

```
go build main.go palindrome.go
```

This will generate a `main` file. To run the program, execute the command

```
./main {size}
```

Where `{size}` is the maximum size of squares/cubes you want the program to check (30,000 is sufficient
to find the solution to the problem).

To run the unit tests, execute

```
go test
```