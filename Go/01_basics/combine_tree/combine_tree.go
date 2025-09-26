package main

/*
 * File: circle_calc.go
 * Author: Franz Zbinden Garcia
 * Course: COTI 4039_vh1
 * Porpuse: x
 */

import "fmt" //para formatear input y output

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

// The main function is the entry point for all Go programs.
func main() {

	fmt.Println(sum(1, 2))

	fmt.Println(minimum(2, 1))

	fmt.Println(combineThree(2, 1, 3))

}
