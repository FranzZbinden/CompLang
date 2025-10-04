/*
 * File: circle_calc_io.go
 * Author: Antonio F. Huertas
 * Course: COTI 4039-LH1
 * Date: 09/29/2025
 * Purpose: This program computes and displays the area and circumference of
 *          the circle with the given radius.
 */

package main

import (
	"fmt"
	"math"
)

// Computes the area of the circle with the given radius.
func circleArea(rad float64) float64 {
	return math.Pi * math.Pow(rad, 2.0)
}

// Computes the circumference of the circle with the given radius.
func circleCircumference(rad float64) float64 {
	return 2.0 * math.Pi * rad
}

// Reads a value and validates it is a number.
func readNumber(prompt string) float64 {
	// TODO Complete this function
	return 0
}

// Reads a number and validates it is non-negative.
func readNonNegative(prompt string) float64 {
	num := readNumber(prompt)
	for num < 0.0 {
		fmt.Println("\tError! The value is a negative number.")
		num = readNumber(prompt)
	}
	return num
}

// Starts the execution of the program.
func main() {
	radius := readNonNegative("Enter the radius:")
	fmt.Printf("The radius of the circle is %.1f\n", radius)

	fmt.Printf("\nIts area is %.1f and its circumference is %.1f\n",
		circleArea(radius), circleCircumference(radius))
}
