/*
 * File: students_eval.go
 * Author: Antonio F. Huertas
 * Course: COTI 4039-LH1
 * Date: 09/24/2025
 * Purpose: This program evaluates a group of students to display their data,
 *          average and grade.
 */

package main

import "fmt"

// Represents a student with id, name, and list of exams.
type student struct {
	id, name string
	exams    []int //exams tipo arreglos
}

// Represents an evaluation consisting of average and grade.
type evaluation struct {
	average float64
	grade   rune //char = rune
}

// Returns the average of the given slice of integers.
func average(slice []int) float64 {
	total := 0
	for _, elem := range slice {
		total += elem
	}
	return float64(total) / float64(len(slice))
}

// Determines the grade corresponding to the given score.
func grade(score float64) rune {
	if score >= 90.0 {
		return 'A'
	} else if score >= 80.0 {
		return 'B'
	} else if score >= 70.0 {
		return 'C'
	} else if score >= 60.0 {
		return 'D'
	}
	return 'F'
}

// Evaluates the given student.
func evaluate(std student) evaluation {
	avg := average(std.exams)
	grd := grade(avg)
	return evaluation{average: avg, grade: grd}
}

// Prints the data and evaluation of the given student.
func printData(std student) {
	fmt.Println("\nId:", std.id)
	fmt.Println("Name:", std.name)
	fmt.Println("Exams:", std.exams)

	eval := evaluate(std)
	fmt.Printf("Average: %.1f\n", eval.average)
	fmt.Printf("Grade: %c\n", eval.grade)
}

// Starts the execution of the program.
func main() {
	students := []student{
		{id: "1111", name: "John Doe", exams: []int{80, 85, 93}},
		{id: "2222", name: "Jane Doe", exams: []int{90, 85, 96}},
		{id: "3333", name: "Joe Smith", exams: []int{81, 76, 80}},
	}

	fmt.Println("These are the students...")
	for _, student := range students {
		printData(student)
	}
}
