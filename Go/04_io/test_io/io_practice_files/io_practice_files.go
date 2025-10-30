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
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	return fmt.Sprintf("width: %.2f, Depth: %.2f, Height: %.2f\n", c.sqr.width, c.sqr.depth, c.height)
}

func (c *cube) scale_cub(scalar float64) {
	c.height *= scalar
	c.sqr.depth *= scalar
	c.sqr.width *= scalar
}

func str_to_cube(line string) cube {
	var c cube

	data := strings.Split(line, "|")
	depth, err := strconv.ParseFloat(data[0], 64)
	if err != nil {
		panic("Error converting from String to Float")
	}
	width, err := strconv.ParseFloat(data[1], 64)
	if err != nil {
		panic("Error converting from String to Float")
	}
	height, err := strconv.ParseFloat(data[2], 64)
	if err != nil {
		panic("Error converting from String to Float")
	}
	c.sqr.width = width
	c.sqr.depth = depth
	c.height = height
	return c
}

func main() {

	// Using manual inputs from console
	// input_width := readPositive("Enter the width of a cube: ")
	// input_depth := readPositive("Enter the depth of a cube: ")
	// input_height := readPositive("Enter the height of a cube: ")

	// in_file_name := "Go/04_io/io_practice_files/measurements.txt"
	// out_file_name := "cubes_strings.txt"

	// in_file, err := os.Open(in_file_name)
	// if err != nil {
	// 	panic("error opening file " + in_file_name)
	// }
	// defer in_file.Close()

	// out_file, err := os.Create(out_file_name)
	// if err != nil {
	// 	panic("Error creating " + out_file_name)
	// }
	// defer out_file.Close()

	// scanner := bufio.NewScanner(in_file)

	// for scanner.Scan() {
	// 	str_line := scanner.Text()

	// 	cub := str_to_cube(str_line)
	// 	fmt.Println(cub.toString())

	in_file_name := "Go/04_io/io_practice_files/measurements.txt"
	out_file_name := "measurements_out.txt"

	in_file, err := os.Open(in_file_name)
	if err != nil {
		panic("Error opening File")
	}
	defer in_file.Close()

	out_file, err := os.Create(out_file_name)
	if err != nil {
		panic("Error creating file " + out_file_name)
	}
	defer out_file.Close()

	scanner := bufio.NewScanner(in_file)

	for scanner.Scan() {
		line := scanner.Text()
		cube := str_to_cube(line)
		fmt.Println(cube)
		out_file.WriteString(cube.toString())
	}

	// input_width
	// input_depth
	// input_height

	// sqr := create_sqr(input_width, input_depth)

	// cube := create_cbe(sqr, input_height)

	// fmt.Println(input_width, input_depth, input_height)
	// fmt.Println()
	// fmt.Println(cube.toString())

	// fmt.Println("The volume of the cube is: ", calc_vol_cube(cube))

	// scalar := 2.0
	// cube.scale_cub(scalar)
	// fmt.Println("The new cube scaled to ", scalar)
	// fmt.Println(cube.toString())

}
