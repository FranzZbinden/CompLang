/*
 * File: combine_three.go
 * Author: Franz Zbinden
 * Date: 09/10/2025
 * Purpose: This program computes and displays the result of combining three
 *          numbers using the given two-argument function.
 */

package main

import "fmt"

// Computes the sum of two numbers.
func sum(num1, num2 int) int {
	return num1 + num2
}

// Computes the minimum of two numbers.
func minimum(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

// Combines three numbers using the given two-argument function.
func combineThree(num1, num2, num3 int, combineTwo func(int, int) int) int {
	temp := combineTwo(num1, num2)
	return combineTwo(temp, num3)
}

// Starts the execution of the program.
func main() {
	number1, number2, number3 := 4, 2, 6
	fmt.Printf("The numbers are %d, %d, and %d\n", number1, number2, number3)

	fmt.Println("\nTheir sum is", combineThree(number1, number2, number3, sum))
	fmt.Println("Their minimum is", combineThree(number1, number2, number3, minimum))

	// Computes the minimum of two numbers.
	maximum := func(n1, n2 int) int {
		if n1 > n2 {
			return n1
		}
		return n2
	}
	fmt.Println("Their maximum is", combineThree(number1, number2, number3, maximum))

	fmt.Println("Their product is", combineThree(number1, number2, number3,
		func(n1, n2 int) int {
			return n1 * n2
		}))
}
