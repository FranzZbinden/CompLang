/*
 * File: fibonacci.go
 * Author: Franz Zbinden Garcia
 * Course: COTI 4039-VH1
 * Purpose: A file compleating the first part of the Assigment#2
 */

package main

import (
	"fmt"
)

// Fibonacci using iteration
func fibonacciIter(num int) int {
	a, b := 0, 1

	if num == 0 {
		return 0
	}
	for cnt := 0; num > cnt; cnt++ {
		temp := a
		a = b
		b = temp + b
	}
	return a
}

func fibonacciiter(num int) int {
	a := 0
	b := 1

	if num == 0 {
		return 0
	}
	for i := 1; i < num; i++ {
		temp := a
		a = b
		b = temp + b
	}
	return a
}

func fibonaccirec(val int) int {
	if val == 0 {
		return 0
	}
	if val == 1 {
		return 1
	}
	return fibonaccirec(val-1) + fibonaccirec(val-2)

}

// Fibonacci using recursion
func fibonacciRec(num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}
	return fibonacciRec(num-1) + fibonacciRec(num-2)
}

// Generates generates Fibonacci numbers using an anonymous function.
func fibonacciGen() func() int {
	a, b := 0, 1
	return func() int {
		next := a
		a, b = b, a+b
		return next
	}
}

func main() {

	fmt.Printf("Using iteration, the 7-th term in the Fibonacci sequence is %d\n", fibonacciIter(7))
	fmt.Println()
	fmt.Printf("Using recursion, the 7-th term in the Fibonacci sequence is %d\n", fibonacciRec(7))
	fmt.Println()
	// Create the Fibonacci generator
	fib := fibonacciGen()

	fmt.Println("Generating the first 10 terms in the Fibonacci sequence…")
	for i := 0; i < 11; i++ {
		fmt.Printf("%d\n", fib())
	}
	fmt.Println()

}
