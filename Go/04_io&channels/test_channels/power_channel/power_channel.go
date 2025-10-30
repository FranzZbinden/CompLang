/*
 * File: factorial_chan.go
 * Author: Franz Zbinden
 * Course: COTI 4039-LH1
 * Date: 10/06/2025
 * Purpose: This program uses channels to concurrently display the power
 *          of some non-negatives integers.
 */

package main

import "fmt"

// Represents a number-power pair.
type pair struct {
	number int
	power  int
}

// 2^num
func powerOfTwoToChan(num int, ch chan<- pair) {
	result := 1

	for range num {
		result = result * 2
	}
	ch <- pair{number: num, power: result}
}

// Returns a channel with the power sequence up to the given limit.
func powerSeq(limit int) <-chan int {
	result := 1
	ch := make(chan int)

	go func() {
		for range limit {
			ch <- result
			result = result * 2
		}
		close(ch)
	}()

	return ch
}

// Sum all results received from a channel of integers.
func sumPowers(ch <-chan int) int {
	sum := 0
	for num := range ch {
		sum += num
	}
	return sum
}

//Receive integers from one channel, send only even ones to another.

func filterEvenPowers(ch <-chan int) <-chan int {
	even_chan := make(chan int)
	go func() {
		for val := range ch {
			if val%2 == 0 {
				even_chan <- val
			}
		}
		close(even_chan)
	}()
	return even_chan
}

// Receive power values from a channel and return the largest. only positives

func maxPower(ch <-chan int) int {
	max := 0
	for val := range ch {
		if val > max {
			max = val
		}
	}
	return max
}

// Sends each input value multiplied by 2 to the output channel.
func doubleValues(ch <-chan int) <-chan int {
	ch_out := make(chan int)

	go func() {
		for val := range ch {
			ch_out <- val * 2
		}
		close(ch_out)
	}()
	return ch_out

}

// Starts the execution of the program.
func main() {
	numbers := []int{7, 10, 0, 12, 1, 5}

	// fmt.Println("Concurrent, possibly parallel, execution:")
	// fmt.Println("Computing the factorial for a group of numbers...")

	pairChan := make(chan pair, len(numbers))
	for _, number := range numbers {
		go powerOfTwoToChan(number, pairChan)
	}

	for range len(numbers) {
		pair := <-pairChan
		fmt.Printf("power of %d is %d\n", pair.number, pair.power)
	}

	limit := 10
	fmt.Println("\nConcurrent, but not parallel, execution:")
	fmt.Println("The power sequence from 1 up to", limit)

	for fact := range powerSeq(limit) {
		fmt.Println(fact)
	}
}
