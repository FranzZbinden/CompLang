/*
 * File: circles_calc.go
 * Author: Franz Zbinden
 * Date: 09/24/2025
 * Purpose: This program computes and displays the area and circumference of
 *          the given circles.
 */

package main

import (
	"fmt"
	"math"
)

// Represents a 2D point with x- and y-coordinates.
type point struct {
	x, y float64
}

// Represents a circle with location and radius.
type circle struct {
	location point
	radius   float64
}

// Computes the area of the given circle.
func circleArea(c circle) float64 {
	return math.Pi * math.Pow(c.radius, 2.0)
}

// Computes the circumference of the given circle.
func circleCircumference(c circle) float64 {
	return 2.0 * math.Pi * c.radius
}

// Scales the given circle by the given factor.
func scaleCircle(c *circle, factor float64) {
	c.radius *= factor
}

// Computes the area of the circle.
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2.0)
}

// Computes the circumference of the circle.
func (c circle) circumference() float64 {
	return 2.0 * math.Pi * c.radius
}

// Scales the circle by the given factor.
func (c *circle) scale(factor float64) {
	c.radius *= factor
}

// Starts the execution of the program.
func main() {
	c1 := circle{location: point{3.0, 4.0}, radius: 5.0}

	fmt.Printf("The location of the first circle is (%.1f, %.1f)\n",
		c1.location.x, c1.location.y)
	fmt.Printf("Its radius is %.1f\n", c1.radius)
	fmt.Printf("Its area is %.1f and its circumference is %.1f\n",
		circleArea(c1), circleCircumference(c1))

	scaleCircle(&c1, 2.0)
	fmt.Printf("\nAfter scaling by 2.0, its radius is now %.1f\n", c1.radius)
	fmt.Printf("Its area is %.1f and its circumference is %.1f\n",
		circleArea(c1), circleCircumference(c1))

	c2 := circle{location: point{9.0, 8.0}, radius: 7.0}

	fmt.Printf("\nThe location of the second circle is (%.1f, %.1f)\n",
		c2.location.x, c2.location.y)
	fmt.Printf("Its radius is %.1f\n", c2.radius)
	fmt.Printf("Its area is %.1f and its circumference is %.1f\n",
		c2.area(), c2.circumference())

	c2.scale(2.0)
	fmt.Printf("\nAfter scaling by 2.0, its radius is now %.1f\n", c2.radius)
	fmt.Printf("Its area is %.1f and its circumference is %.1f\n",
		c2.area(), c2.circumference())
}
