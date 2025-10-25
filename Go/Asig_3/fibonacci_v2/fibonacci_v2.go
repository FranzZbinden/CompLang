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

// func fibonacciFiniteSeq(limit int) iter.Seq[int]
// 	var fibos [limit]int
// 	func(){
// 		a, b := 0, 1

// 		if num == 0 {
// 			return 0
// 		}
// 		for cnt := 0; num > cnt; cnt++ {
// 			temp := a
// 			a = b
// 			b = temp + b
// 		}
// 		return a
// 	}()

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

	var value string

	fmt.Print("Enter a positive limit: ")
	fmt.Scan(&value)
	limit, err := strconv.Atoi(value)
	if err != nil {
		panic("Error entering data")
	}
	print(limit)
}
