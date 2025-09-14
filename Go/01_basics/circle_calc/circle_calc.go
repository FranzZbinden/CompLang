package circlecalc

/*
 * File: circle_calc.go
 * Author: Franz Zbinden Garcia
 * Course: COTI 4039_vh1
 * Porpuse: x
 */

import (
	"fmt"
	"math"
)

// Computes the area of the circle with the given radius.
func circleArea(rad float64) float64 { //en tipo de dato va al final de la variable
	return math.Pi * math.Pow(rad, 2.0)

}

// Computes the circumference of the circle with the given radius.
func circleCircumference(rad float64) float64 { //en tipo de dato va al final de la variable
	return 2.0 * math.Pi * rad
}

func main() {
	radius := 5.0

	fmt.Printf("The radius of the circle is %.1f\n", radius) //Printf: se usa para imprimir floats

	fmt.Printf("\nIts area is %.1f, and its circumference is, %.1\n", circleArea(radius), circleCircumference(radius))

	fmt.Println("Hello, World!") // la primera letra de una funcion indica que la funcion es publica (ejemplo: Println)

}
