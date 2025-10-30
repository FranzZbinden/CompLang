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

type square struct {
	width float64
	depth float64
}

type cube struct {
	sqr    square
	height float64
}

func create_sqr(width float64, depth float64) square {
	return square{width: width, depth: depth}
}

func create_cbe(sqr square, height float64) cube {
	return cube{sqr: sqr, height: height}
}

func calc_vol_cube(c cube) float64 {
	return c.height * c.sqr.depth * c.sqr.width
}

// Reads a value and validates it is a number.
func readNumber(prompt string) float64 {
	var in float64
	fmt.Print(prompt)
	_, err := fmt.Scanln(&in)
	if err != nil {
		panic("Error entering values")
	}
	return in
}

// Reads a number and validates it is positive.
func readPositive(prompt string) float64 {
	num := readNumber(prompt)
	if num < 0 {
		num = readNumber("Number must be positive float")
	}
	return num
}

func (c cube) toString() string {
	return fmt.Sprintf("width: %.2f, Depth: %.2f, Height: %.2f", c.sqr.width, c.sqr.depth, c.height)
}

func main() {
	input_width := readPositive("Enter the width of a cube: ")
	input_depth := readPositive("Enter the depth of a cube: ")
	input_height := readPositive("Enter the height of a cube: ")

	sqr := create_sqr(input_width, input_depth)

	cube := create_cbe(sqr, input_height)

	fmt.Println(input_width, input_depth, input_height)
	fmt.Println()
	fmt.Println(cube.toString())

	fmt.Println("The volume of the cube is: ", calc_vol_cube(cube))

}
