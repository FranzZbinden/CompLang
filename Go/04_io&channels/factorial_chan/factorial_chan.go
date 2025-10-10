/*
 * File: factorial_chan.go
 * Author: Antonio F. Huertas
 * Course: COTI 4039-LH1
 * Date: 10/06/2025
 * Purpose: This program uses channels to concurrently display the factorial
 *          of some non-negatives integers.
 */

package main

import (
	"fmt"
)

// Represents a number-factorial pair.
type pair struct {
	number    int
	factorial int
}

// Computes the factorial of the given number and puts the number-factorial pair
// into the given channel.
// concurrencia es una funcion no determinista
// concurrencia con pararelismo
func factorialToChan(num int, ch chan<- pair) { //ch es un canal en la que puedo colocar pairs
	prod := 1
	for cnt := 1; cnt <= num; cnt++ {
		prod *= cnt
	}

	ch <- pair{number: num, factorial: prod} //saca del canal y lo pone en ch
}

// Returns a channel with the factorial sequence up to the given limit.
func factorialSeq(limit int) <-chan int { // concurrencia pero no paralelismo
	prod := 1
	ch := make(chan int) //crea un canal sin tamaño , sin buffer

	go func() { // esto corre, luego pausa, main lo muestra, main pausa, func genera prod y asi sicesivamente
		defer close(ch)
		for cnt := 1; cnt <= limit; cnt++ {
			prod *= cnt
			ch <- prod
		}
	}() //el () significa que se esta llamando ahi mismo
	return ch
}

// Starts the execution of the program.
func main() {
	numbers := []int{0, 1, 5, 7, 10, 12}

	fmt.Println("Concurrent, possibly parallel, execution:")
	fmt.Println("Computing the factorial for a group of numbers...")

	pairChan := make(chan pair, len(numbers))
	for _, number := range numbers {
		go factorialToChan(number, pairChan) // go routines means
	}

	for range len(numbers) {
		pair := <-pairChan //sacando del canal par aponer en pair
		fmt.Printf("Factorial of %d is %d\n", pair.number, pair.factorial)
	}

	limit := 10
	fmt.Println("\nConcurrent, but not parallel, execution:")
	fmt.Println("The factorial sequence from 1 up to", limit)

	for fact := range factorialSeq(limit) { // secuencialmente esto corre facotrial sec que produce resultado e imprime antes de producir el proximo valor
		fmt.Println(fact)
	}
}
