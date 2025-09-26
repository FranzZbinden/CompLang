/*
 * File: circle_calc.go
 * Author: Franz Zbinden Garcia
 * Course: COTI 4039-LH1
 * Date: 09/08/2025
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
	return math.Pi * math.Pow(rad, 2.0) //math.Pow: eleva un numero a una potencia 2.0
}

// Computes the circumference of the circle with the given radius.
func circleCircum(rad float64) float64 { //el tipo de dato va al final de la variable, no antes
	return 2.0 * math.Pi * rad
}

// Starts the execution of the program.
func main() {
	radius := 5.0
	fmt.Printf("The radius of the circle is %.1f\n", radius) //Printf: se usa para imprimir floats

	fmt.Printf("\nIts area is %.1f and its circumference is %.1f\n",
		circleArea(radius), circleCircum(radius))
}
