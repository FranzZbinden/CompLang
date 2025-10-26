/*
 * File: fibonacci_v2.go
 * Author: Franz Zbinden García
 * Course: COTI 4039-LH1
 * Date: 10/16/2025
 * Purpose: This program uses channels to concurrently display the fibonacci
 *          of some non-negatives integers.
 */

package main

import (
	"fmt"
	"iter"
	"strconv"
)

type pair struct {
	num  int
	fibo int
}

// Returns a fibonacci pair acording to the given number
func fibonacciToChan(num int, ch chan<- pair) {
	a, b := 0, 1

	if num == 0 {
		ch <- pair{num: num, fibo: 0}
		return
	}
	for cnt := 0; num > cnt; cnt++ {
		temp := a
		a = b
		b = temp + b
	}
	ch <- pair{num: num, fibo: a}
}

// returns an iterator over the the Fibonacci sequence up to the given limit
func fibonacciFiniteSeq(limit int) iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1
		for cnt := 1; cnt <= +limit; cnt++ {
			if !yield(b) { // yield first so it starts at 1st fibo num
				return
			}
			temp := a
			a = b
			b = temp + b
		}
	}
}

// reads and returns the limit, if error: calls again for correct answear
func read_number(prompt string) int {
	var value_in string
	fmt.Print(prompt)
	fmt.Scan(&value_in)
	limit, err := strconv.Atoi(value_in)
	if err != nil {
		return read_number("Number must be a positive integer: ")
	}
	return limit
}

func main() {
	nums := []int{7, 8, 6}

	ch := make(chan pair, len(nums))

	// Calls the go routines
	for _, num := range nums {
		go fibonacciToChan(num, ch)
	}

	fmt.Println("The slice is ", nums)
	fmt.Println("Computing the nth Fibonacci number for each element…")

	// Recives & prints the fibonacci pairs
	for range len(nums) {
		pairFibo := <-ch
		fmt.Printf("Fibonacci of %d is %d\n", pairFibo.num, pairFibo.fibo)
	}

	// sets the limit fron the input of the user
	limit := read_number("Enter a positive limit:")

	// prints the fibonacci sequence
	fmt.Println("\nThe fibonacci sequence from 1 up to", limit)
	index := 0
	for fact := range fibonacciFiniteSeq(limit) {
		index += 1
		fmt.Println("The fibonacci of", index, "is", fact)
	}
}
