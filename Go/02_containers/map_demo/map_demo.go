/*
 * File: map_demo.go
 * Author: Antonio F. Huertas
 * Course: COTI 4039-LH1
 * Date: 09/22/2025
 * Purpose: This program demonstrates how to define and use maps.
 */

package main

import "fmt"

// Prints the map of courses as a table.
func printData(courseMap map[string]int) {
	fmt.Println("Title\t  Credits")
	fmt.Println("-----\t  -------")

	for title, credits := range courseMap {
		fmt.Println(title, "\t   ", credits)
	}
	fmt.Println("There are", len(courseMap), "courses")
}

// Entry point of the program.
func main() {
	courses := map[string]int{
		"COTI3101": 4, "SICI3205": 3, "SICI4036": 3, "COTI4039": 3,
	}

	fmt.Println("The map of courses to credits is", courses)
	fmt.Println("Its length is", len(courses))

	fmt.Println("\nThese are the courses...")
	printData(courses)

	newCourse := "COTI3102"
	courses[newCourse] = 3

	fmt.Println("\nAfter adding", newCourse, "these are the courses...")
	printData(courses)

	courses[newCourse] = 4

	fmt.Println("\nAfter modifying", newCourse, "these are the courses...")
	printData(courses)

	credits, found := courses[newCourse]

	fmt.Print("\nThe course with code ", newCourse, " was ")
	if found {
		fmt.Println("found and has", credits, "credits")
	} else {
		fmt.Println("not found")
	}

	delete(courses, newCourse)
	fmt.Println("\nAfter deleting", newCourse, "these are the courses...")
	printData(courses)

	clear(courses)
	fmt.Println("\nAfter clearing, the map of courses to credits is", courses)
	fmt.Println("Its length is", len(courses))

	words := make(map[string]string)
	words["Hello"] = "Hola"
	words["Good-bye"] = "AdiÃ³s"
	words["Today"] = "Hoy"
	words["Yesterday"] = "Ayer"
	words["Tomorrow"] = "MaÃ±ana"

	fmt.Println("\nThese are the words...")
	for english, spanish := range words {
		fmt.Println("English:", english, "=> Spanish:", spanish)
	}
	fmt.Println("The number of items is", len(words))
}
