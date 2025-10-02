/*
 * File: fibonacci.go
 * Author: Franz Zbinden Garcia
 * Course: COTI 4039-VH1
 * Purpose: A file compleating the first part of the Assigment#2
 */

package main

import (
	"cmp"
	"fmt"
)

// Search for a target in a slice and return the position
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

// Search for a target in a sorted slice and return the position
func binarySearch[E cmp.Ordered](slice []E, target E) (int, bool) {
	if len(slice) == 0 {
		return 0, false
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
		return middle + 1 + index, found
	}
}

// Returns true if a slice is sorted in ascending order
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

// Sorts a slice in place using insertion sort algorithm
func insertionSort[E cmp.Ordered](slice []E) {

	for i := 1; i < len(slice); i++ {
		key := slice[i]
		j := i - 1

		for j >= 0 && slice[j] > key {
			slice[j+1] = slice[j]
			j--
		}

		slice[j+1] = key
	}
}

// Eeturns the minimum and maximum of a variable number of values
func extrema[E cmp.Ordered](values ...E) (min, max E) {
	if len(values) == 0 {
		panic("No values provided!!")
	}

	min = values[0]
	max = values[0]

	for index := 1; index < len(values); index++ {
		if values[index] < min {
			min = values[index]
		}
		if values[index] > max {
			max = values[index]
		}
	}
	return min, max
}

func isSortedNums[E cmp.Ordered](slice []E) string {
	if isSorted(slice) {
		return "\nThe slice is sorted"
	}
	return "\nThe slice is not sorted"
}

func main() {

	nums := []int{10, -1, 0, 3, 28, -30, -7, 41, -6, 39}

	fmt.Printf("The slice is %v\n", nums)

	max, min := extrema(10, -1, 0, 3, 28, -30, -7, 41, -6, 39)
	fmt.Printf("The extrema are %d and %d\n", min, max)

	fmt.Println()

	key := -6
	index, found := binarySearch(nums, key)
	fmt.Printf("Does it contain -6? %v\n", found)
	fmt.Printf("Key %d was found at index #%d\n", key, index)

	key2 := 20
	_, found2 := binarySearch(nums, key2)
	fmt.Printf("Does it contain 20? %v", found2)

	fmt.Println()

	fmt.Println(isSortedNums(nums))
	insertionSort(nums)
	fmt.Printf("The slice in ascending order is %v", nums)
	fmt.Println(isSortedNums(nums))

	fmt.Println()

	index3, found3 := binarySearch(nums, key)
	fmt.Printf("Does it contain %d? %v\n", key, found3)
	fmt.Printf("Key %d was found at index #%d\n", key, index3)

	index4, found4 := binarySearch(nums, 20)
	fmt.Printf("Does it contain %d? %v\n", 20, found4)
	fmt.Printf("Key %d should be inserted at index #%d\n", 20, index4)

}
