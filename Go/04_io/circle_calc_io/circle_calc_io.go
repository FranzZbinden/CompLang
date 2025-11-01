/*
 * File: circle_calc_io.go
 * Author: Franz Zbinden
 * Date: 09/29/2025
 * Purpose: This program computes and displays the area and circumference of
 *          the circle with the given radius.
 */

package main

import (
	"fmt"
	"math"
	"strconv"
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
	for {
		var value string
		fmt.Print("\n" + prompt + " ")
		fmt.Scanln(&value)

		num, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Println("\tError! The value is not a number")
		} else {
			return num
		}
	}
}

// Reads a number and validates it is positive.
func readPositive(prompt string) float64 {
	num := readNumber(prompt)
	for num <= 0.0 {
		fmt.Println("\tError! The value is not a positive number")
		num = readNumber(prompt)
	}
	return num
}

// Starts the execution of the program.
func main() {
	radius := readPositive("Enter the radius:")
	fmt.Printf("The radius of the circle is %.1f\n", radius)

	fmt.Printf("\nIts area is %.1f and its circumference is %.1f\n",
		circleArea(radius), circleCircumference(radius))

	// ===========================================================

	rad := readNumber("Enter the radius of a circle:")
	fmt.Printf("\nThe radius of the circle is %.2f\n", rad)

	circle_area, circle_circumference := circleArea(rad), circleCircumference(rad)
	fmt.Printf("The area of the circle is: %.2f, and the circumference is: %.2f\n",
		circle_area, circle_circumference)
}
