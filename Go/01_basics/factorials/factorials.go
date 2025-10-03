/*
 * File: factorials.go
 * Author: Franz Zbinden
 * Course: COTI 4039-LH1
 * Date: 09/15/2025
 * Purpose: This program computes and displays the factorial of a non-negative integer
 *          using various techniques.
 */

package main

import "fmt"

// Computes the factorial of an integer using a simple loop.
func factorial(num int) int { //like a while loop, in go thers no while loop
	prod := 1
	for num > 0 {
		prod *= num // == than prod = prod * num
		num--
	}
	return prod
}

// Computes the factorial of an integer using a three-part for loop.
func factorial2(num int) int {
	prod := 1
	for cnt := 1; cnt <= num; cnt++ { //this is a traditional for loop
		prod *= cnt
	}
	return prod
}

// Computes the factorial of an integer using for-range loop.
func factorial3(num int) int {
	prod := 1
	for cnt := range num {
		prod *= cnt + 1
	}
	return prod
}

// Computes the factorial of an integer using recursion.
func factorialRec(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorialRec(num-1)
}

func factorialRec2(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorialRec2(num-1)
}

// Generates each value in the factorial sequence by returning a closure.
func factorialGen() func() int {
	prod, cnt := 1, 1
	return func() int {
		prod *= cnt
		cnt++
		return prod
	}
}




//a function that when called, it sets a recording of how many times it has benn called
func numGen()func() int{
	count := 0
	returnfunc()int{
		return count += 1
	}
}













// Starts the execution of the program.
func main() {
	number := 5
	fmt.Println("The number is", number)

	fmt.Println("\nIts factorial is...")
	fmt.Println("\tUsing a simple loop:", factorial(number))
	fmt.Println("\tUsing a three-part for loop:", factorial2(number))
	fmt.Println("\tUsing a for-range loop:", factorial3(number))
	fmt.Println("\tUsing recursion:", factorialRec(number))

	limit := 10
	fmt.Println("\nThe factorial sequence from 1 up to", limit)

	nextValue := factorialGen()
	for range limit {
		fmt.Println(nextValue())
	}
}
