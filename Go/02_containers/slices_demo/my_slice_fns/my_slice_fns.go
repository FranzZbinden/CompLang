/*
 * File: my_slice_fns.go
 * Author: Antonio F. Huertas
 * Course: COTI 4039-LH1
 * Date: 09/17/2025, 09/22/2025
 * Purpose: This program demonstrates how to define some slice functions.
 */

package main

import (
	"cmp"
	"fmt"
)

// Computes the minimum element of a slice.
func minimum[E cmp.Ordered](slice []E) E {
	if len(slice) == 0 {
		panic("empty slice")
	}
	min := slice[0]
	for index := 1; index < len(slice); index++ {
		if slice[index] < min {
			min = slice[index]
		}
	}
	return min
}

// Computes the maximum element of a slice.
func maximum[E cmp.Ordered](slice []E) E {
	if len(slice) == 0 {
		panic("empty slice")
	}
	max := slice[0]
	for index := 1; index < len(slice); index++ {
		if slice[index] > max {
			max = slice[index]
		}
	}
	return max
}

// Determines whether a slice contains the target.
func contains[E comparable](slice []E, target E) bool {
	for _, elem := range slice {
		if elem == target {
			return true
		}
	}
	return false
}

// Searches a sorted slice for a target and returns its index if found, or
// -1 if not found.
func index[E comparable](slice []E, target E) int {
	for index, elem := range slice {
		if elem == target {
			return index
		}
	}
	return -1
}

// Searches a sorted slice for a target and returns its index and true if found, or
// -1 and false if not found.
// Precondition: The slice is sorted in ascending order.
func binarySearch[E cmp.Ordered](slice []E, target E) (int, bool) {
	low, high := 0, len(slice)-1
	for low <= high {
		middle := (low + high) / 2
		if slice[middle] < target {
			low = middle + 1
		} else if slice[middle] > target {
			high = middle - 1
		} else {
			return middle, true
		}
	}
	return low, false
}

// Sorts a slice in ascending order.
func quickSort[E cmp.Ordered](slice []E) {
	if len(slice) <= 1 {
		return
	}
	low, high := 0, len(slice)-1
	for index := range slice {
		if slice[index] < slice[high] {
			slice[index], slice[low] = slice[low], slice[index]
			low++
		}
	}
	slice[low], slice[high] = slice[high], slice[low]
	quickSort(slice[:low])
	quickSort(slice[low+1:])
}

// Reverses the elements of a slice.
func reverse[E any](slice []E) {
	low, high := 0, len(slice)-1
	for low < high {
		slice[low], slice[high] = slice[high], slice[low]
		low++
		high--
	}
}

// Starts the execution of the program.
func main() {
	numbers := []int{30, 10, 50, 40, 20}

	fmt.Println("The slice of numbers is", numbers)
	fmt.Println("Its length is", len(numbers))
	fmt.Println("Its capacity is", cap(numbers))

	fmt.Println("\nIts minimum element is", minimum(numbers))
	fmt.Println("Its maximum element is", maximum(numbers))

	fmt.Println("\nDoes it contain 20?", contains(numbers, 20))
	fmt.Println("Index of 20?", index(numbers, 20))

	quickSort(numbers)
	fmt.Println("\nThe sorted slice is", numbers)

	index, found := binarySearch(numbers, 20)
	fmt.Println("\nDoes it contain 20?", found)
	fmt.Println("Index of 20?", index)

	index, found = binarySearch(numbers, 25)
	fmt.Println("\nDoes it contain 25?", found)
	fmt.Println("Index where 25 should be placed?", index)

	reverse(numbers)
	fmt.Println("\nThe reversed slice is", numbers)

	fmt.Println("\nThe sum of the elements is", sum(numbers...))
	fmt.Println("The sum of the values is", sum(30, 10, 50, 40, 20))
}

// Returns the sum of a variable number of values.
func sum(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}
