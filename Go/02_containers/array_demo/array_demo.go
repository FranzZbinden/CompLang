/*
 * File: array_demo.go
 * Author: Franz Zbinden
 * Date: 09/15/2025
 * Purpose: This program demonstrates how to define and use arrays.
 */

package main

import "fmt"

// Sums the elements of an array.
func sum(arr [5]int) int {
	total := 0
	for _, elem := range arr {
		total += elem
	}
	return total
}

// Searches an array for a target and returns its index and true if found or
// -1 and false if not found.
func search(arr [5]int, target int) (int, bool) {
	for index, elem := range arr {
		if elem == target {
			return index, true
		}
	}
	return -1, false
}

func powNums(arr [5]int) [5]int {
	var arr1 [5]int
	for index, val := range arr {
		arr1[index] = val * val
	}
	return arr1
}

// Starts the execution of the program.
func main() {
	numbers := [...]int{30, 10, 50, 40, 20}

	fmt.Println("The array of powers of 30, 10, 50, 40, 20 is", powNums(numbers))

	fmt.Println("The array of numbers is", numbers)

	fmt.Println("\nThese are the elements, one per line: ")
	for index, number := range numbers {
		fmt.Println(index, ":", number)
	}

	fmt.Println("\nIts length is", len(numbers))
	fmt.Println("Its capacity is", cap(numbers))
	fmt.Println("The sum of its elements is", sum(numbers))

	index, found := search(numbers, 20)
	fmt.Println("\nDoes it contain 20?", found)
	fmt.Println("Index of 20?", index)
	fmt.Println("The element at index #4 is", numbers[4])

	var evenSquares [10]int
	index = 0
	for num := 1; num <= 20; num++ {
		if num%2 == 0 {
			evenSquares[index] = num * num
			index++
		}
	}
	fmt.Println("\nThe first ten even squares are", evenSquares)
	fmt.Println("Its four middle elements are", evenSquares[3:7])
}
