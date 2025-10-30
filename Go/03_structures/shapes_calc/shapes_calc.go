/*
 * File: shapes_calc.go
 * Author: Franz Zbinden
 * Course: COTI 4039-LH1
 * Date: 09/24/2025
 * Purpose: This program displays the data, area, and perimeter for each element
 *          in a list of shapes.
 */

package main

import (
	"fmt"
	"math"
)

// iota asigna un valor a todas las variables en seuencia (1,2,3,4,5)
// Represents the set of available colors.
type color int

const (
	black color = iota
	white
	red
	green
	blue
	yellow
)

// ASignacion para enumeracion de los departamentos
// Returns the string representation of the color.
func (c color) String() string {
	if black > c || c > yellow {
		return "unknown"
	}
	colorName := map[color]string{
		black:  "Black",
		white:  "white",
		red:    "Red",
		green:  "Green",
		blue:   "Blue",
		yellow: "Yellow",
	}
	return colorName[c]

}

// Represents a point with x- and y-coordinates.
type point struct {
	x, y float64
}

// Returns the string representation of the point.
func (p point) String() string {
	return fmt.Sprintf("(%.1f, %.1f)", p.x, p.y) //para aplicar formato
}

// Represents the common information for all shapes.
type shapeInfo struct {
	color    //Crear un campo en shapeInfo llamado color
	location point
}

// Returns the string representation of the shape.
func (si shapeInfo) String() string {
	return fmt.Sprintf("color: %s, location: %s", si.color, si.location)
}

// Interface hace una especia de contrato para las funciones, alfo asi
// Represents the set of methods common to any shape.
type shape interface {
	area() float64
	perimeter() float64
}

// Represents a circle as a shape with a radius.
type circle struct {
	shapeInfo
	radius float64
}

type triangle struct {
	color string
}

// Computes the area of the circle.
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2.0)
}

// Computes the perimeter of the circle.
func (c circle) perimeter() float64 {
	return 2.0 * math.Pi * c.radius
}

// Returns the string representation of the circle.
func (c circle) String() string {
	return fmt.Sprintf("%s, radius: %.1f", c.shapeInfo, c.radius)
}

// Represents a rectangle as a shape with width and height.
type rectangle struct {
	shapeInfo
	width, height float64
}

// Computes the area of the rectangle.
func (r rectangle) area() float64 {
	return r.width * r.height
}

// Computes the perimeter of the rectangle.
func (r rectangle) perimeter() float64 {
	return 2.0 * (r.width + r.height)
}

// Returns the string representation of the rectangle.
func (r rectangle) String() string {
	return fmt.Sprintf("%s, width: %.1f, height: %.1f", r.shapeInfo,
		r.width, r.height)
}

// Prints the data of the given shape.
func printData(shp shape) {
	fmt.Print("\nThe shape is a ")
	// TODO: Complete this function

	fmt.Println("Its data is", shp)
	fmt.Printf("Its area is %.1f\n", shp.area())
	fmt.Printf("Its perimeter is %.1f\n", shp.perimeter())
}

// Starts the execution of the program.
func main() {
	shapes := []shape{
		circle{shapeInfo: shapeInfo{color: green, location: point{1.0, 2.0}},
			radius: 5.0},
		rectangle{shapeInfo: shapeInfo{color: yellow, location: point{8.0, 9.0}},
			width: 5.0, height: 6.0},
		circle{shapeInfo: shapeInfo{color: blue, location: point{7.0, 3.0}},
			radius: 4.0},
	}

	fmt.Println("These are the shapes...")
	for _, shape := range shapes {
		printData(shape)
	}
}
