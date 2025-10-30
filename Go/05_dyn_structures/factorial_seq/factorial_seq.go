/*
 * File: factorial_seq.go
 * Author: Franz Zbinden
 * Course: COTI 4039-LH1
 * Date: 10/06/2025
 * Purpose: This program computes and displays the factorial sequence.
 */

package main

import (
	"fmt"
	"iter"
)

// Returns an iterator over the factorial sequence up to the given limit.
func factorialFiniteSeq(limit int) iter.Seq[int] {
	return func(yield func(int) bool) {
		prod := 1
		for cnt := 1; cnt <= limit; cnt++ {
			prod *= cnt
			if !yield(prod) {
				return
			}
		}
	}
}

// Returns an iterator over the infinite factorial sequence.
func factorialInfiniteSeq2() iter.Seq[int] {
	return func(yield func(int) bool) {
		prod, cnt := 1, 1
		for {
			prod *= cnt
			cnt++
			if !yield(prod) {
				return
			}
		}
	}
}

func sumFiniteSeq(limit int) iter.Seq[int] {
	return func(yield func(int) bool) {
		sumation := 0
		for cnt := 0; cnt <= limit; cnt++ {
			sumation += cnt
		}
		if !yield(sumation) {
			return
		}
	}
}

// Starts the execution of the program.
func main() {
	limit := 10

	fmt.Println("The factorial sequence from 1 up to", limit)
	for fact := range factorialFiniteSeq(limit) {
		fmt.Println(fact)
	}

	fmt.Println("\nThe first elements of an infinite factorial sequence...")
	for fact := range factorialInfiniteSeq() {
		fmt.Println(fact)
		if fact > 1000000 {
			break
		}
	}
}
