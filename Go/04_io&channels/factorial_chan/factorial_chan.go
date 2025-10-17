/*
 * File: factorial_chan.go
 * Author: Antonio F. Huertas
 * Course: COTI 4039-LH1
 * Date: 10/06/2025
 * Purpose: This program uses channels to concurrently display the factorial
 *          of some non-negatives integers.
 */

package main

import "fmt"

// Represents a number-factorial pair.
type pair struct {
	number    int
	factorial int
}

// Computes the factorial of the given number and puts the number-factorial pair
// into the given channel.
func factorialToChan(num int, ch chan<- pair) {
	prod := 1

	for cnt := 1; cnt <= num; cnt++ {
		prod *= cnt
	}
	ch <- pair{number: num, factorial: prod}
}

// Returns a channel with the factorial sequence up to the given limit.
func factorialSeq(limit int) <-chan int {
	prod := 1
	ch := make(chan int)

	go func() {
		defer close(ch)
		for cnt := 1; cnt <= limit; cnt++ {
			prod *= cnt
			ch <- prod
		}
	}()
	return ch
}

// Starts the execution of the program.
func main() {
	numbers := []int{7, 10, 0, 12, 1, 5}

	// fmt.Println("Concurrent, possibly parallel, execution:")
	// fmt.Println("Computing the factorial for a group of numbers...")

	pairChan := make(chan pair, len(numbers))
	for _, number := range numbers {
		go factorialToChan(number, pairChan)
	}

	for range len(numbers) {
		pair := <-pairChan
		fmt.Printf("Factorial of %d is %d\n", pair.number, pair.factorial)
	}

	limit := 10
	fmt.Println("\nConcurrent, but not parallel, execution:")
	fmt.Println("The factorial sequence from 1 up to", limit)

	for fact := range factorialSeq(limit) {
		fmt.Println(fact)
	}
}
