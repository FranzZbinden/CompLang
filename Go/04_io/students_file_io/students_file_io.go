/*
 * File: students_file_io.go
 * Author: Antonio F. Huertas
 * Course: COTI 4039-LH1
 * Date: 10/01/2025
 * Purpose: This program reads a group of students from a file, evaluates
 *          them, and creates a new file with their average and grade.
 */

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Represents a student with id, name, and a list of exams.
type student struct {
	id, name string
	exams    []int
}

// Represents an evaluation with an average and a grade.
type evaluation struct {
	average float64
	grade   rune
}

// Computes the average of a slice of integers.
func average(slice []int) float64 {
	sum := 0
	for _, elem := range slice {
		sum += elem
	}
	return float64(sum) / float64(len(slice))
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

// Parses a line and returns the corresponding student or error.
func parseStudentLine(line string) (student, error) {
	var empty student

	fields := strings.Split(line, "|")
	if len(fields) != 5 {
		return empty, errors.New("Incorrect format in line: " + line)
	}

	id := strings.TrimSpace(fields[0])
	name := strings.TrimSpace(fields[1])

	exams := []int{}
	for _, field := range fields[2:] {
		field = strings.TrimSpace(field)
		exam, err := strconv.Atoi(field)
		if err != nil {
			return empty, errors.New("Invalid numeric data in line: " + line)
		}
		exams = append(exams, exam)
	}
	return student{id, name, exams}, nil
}

// Returns the evaluation line for the given student.
func newEvaluationLine(std student) string {
	eval := evaluate(std)
	return fmt.Sprintf("%s|%s|%.1f|%c\n", std.id, std.name,
		eval.average, eval.grade)
}

// Starts the execution of the program.
func main() {
	inFilename := "students.txt"
	outFilename := "evaluations.txt"

	inFile, err := os.Open(inFilename)
	if err != nil {
		panic("Error opening input file " + inFilename)
	}
	defer inFile.Close() // close inFile when main finishes

	outFile, err := os.Create(outFilename)
	if err != nil {
		panic("Error opening output file " + outFilename)
	}
	defer outFile.Close() // close outFile when main finishes

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		student, err := parseStudentLine(scanner.Text())
		if err != nil {
			println(err.Error())
			continue
		}
		outFile.WriteString(newEvaluationLine(student))
	}
}
