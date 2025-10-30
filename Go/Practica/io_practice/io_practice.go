/*
 * File: cube_calc_io.go
 * Author: Franz Zbinden
 * Course: COTI 4039-LH1
 * Date: 09/29/2025
 * Purpose: This program computes and displays the area and circumference of
 *          the circle with the given radius.
 */

package main

import (
	"fmt"
)

// Reads a value and validates it is a number.
func readNumber(prompt string) float64 {
	var in_str int
	fmt.Print(prompt)
	in_str, err := fmt.Scanln(&in_str)
	if err != nil {
		panic("Error entering values")
	}

	if readPositive(in) == -1 {
		panic("Number must be positive")
	}
}

// Reads a number and validates it is positive.
func readPositive(prompt string) float64 {
	if 

}

func main() {
	input_width := readNumber("Enter the width of a cube")
	input_depth := readNumber("Enter the depth of a cube")
	input_height := readNumber("Enter the height of a cube")

}
