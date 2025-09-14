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

func minimum(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

func combineThree(num1, num2, num3 int,
	combineTwo func(int, int) int) int {
	temp := combineTwo(num1, num2)
	return combineTwo(temp, num3)
}

func main() {
	number1, number2, number3 := 4, 2, 6
	fmt.Printf("The numbers are %d, %d, and %d\n", number1, number2, number3)

	fmt.Println("\nTheir sum is", combineThree(number1, number2, number3, sum))
	fmt.Println("\nTheir minimum is", combineThree(number1, number2, number3, minimum))
}
