/*
 * File: fibonacci.go
 * Author: Franz Zbinden Garcia
 * Course: COTI 4039-LH1
 * Purpose: A file compleating the first part of the Assigment#2
 */

package main

import (
	"cmp"
	"fmt"
)

func index[E comparable](slice []E, target E) int {

	if len(slice) == 0 {
		return -1
	}

	if slice[0] == target {
		return 0
	}

	indexNum := index(slice[1:], target)

	return indexNum + 1

}

// A function binarySearch that uses recursion to search for a target in a
// sorted slice and return the position of the target and true if the target
// is found or its insertion point and false if the target was not found:
func binarySearch[E cmp.Ordered](slice []E, target E) (int, bool) {
	if len(slice) == 0 {
		return -1, false
	}
	middle := len(slice) / 2

	if slice[middle] == target {
		return middle, true
	}

	//left half
	if target < slice[middle] {
		return binarySearch(slice[:middle], target)
	} else {
		// search right half
		index, found := binarySearch(slice[middle+1:], target)
		if found {
			// adjust index, left + middle
			return middle + 1 + index, found
		}
		return -1, false
	}
}

// function isSorted that returns true if a slice is sorted in ascending order:
func isSorted[E cmp.Ordered](slice []E) bool {

	if len(slice) <= 1 {
		return true
	}
	var sorted bool

	if slice[0] <= slice[1] {
		sorted = isSorted(slice[1:])
	}
	if sorted == true {
		return true
	}
	return false
}

// function insertionSort that sorts a slice in place using insertion sort algorithm:
func insertionSort[E cmp.Ordered](slice []E) {
	// Outer loop: start from index 1 (because a slice of 1 element is already "sorted")
	for i := 1; i < len(slice); i++ {

		// Take the current element we want to insert into the sorted left side
		key := slice[i] // "key" = the element we are trying to put in the right position

		// Start comparing from the element just before key (the left side of the slice)
		j := i - 1

		// Inner loop: shift elements to the right until the right spot for "key" is found
		// Condition means: "while we haven’t gone past the start AND slice[j] is bigger than key"
		for j >= 0 && slice[j] > key {
			// Shift slice[j] one position to the right
			slice[j+1] = slice[j]

			// Move j one step left (to keep checking earlier elements)
			j--
		}

		// When we break out of the loop, j is at the position just before where key belongs
		// So insert key at the correct sorted position
		slice[j+1] = key
	}
}

func main() {

	nums := []int{10, 20, 30, 40, 50}

	// Recursive tests
	fmt.Println("Recursive index of 30:", index(nums, 30)) // expect 2
	fmt.Println("Recursive index of 50:", index(nums, 50)) // expect 4
	fmt.Println("Recursive index of 99:", index(nums, 99)) // expect -1

	data := []int{10, 20, 30, 40, 50}

	fmt.Println(binarySearch(data, 30)) // (2, true)
	fmt.Println(binarySearch(data, 10)) // (0, true)
	fmt.Println(binarySearch(data, 99)) // (-1, false)

	fmt.Println()

	fmt.Println(isSorted(data))
}
