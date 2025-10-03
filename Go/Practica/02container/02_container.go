/*
 * File: payroll.go
 * Author: Franz Zbinden Garcia
 * Course: COTI 4039-VH1
 * Date: 09/30/2025
 * Purpose: Part 3 of the Assigment.
 */

package main

import (
	"cmp"
	"fmt"
)

func index(slice []int, target int) int {
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
	mid := len(slice) / 2

	if slice[mid] == target {
		return mid, true
	}
	if slice[mid] > target {
		return binarySearch(slice[:mid], target)
	} else {
		indexNum, found := binarySearch(slice[mid:], target)
		return indexNum + mid + 1, found
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
	if sorted {
		return true
	}
	return false
}

func power(slice []int, power int) {
	for i := 0; i < len(slice); i++ {
		result := slice[i]
		for j := 0; j < power; j++ {
			result = result * slice[i]
		}
		slice[i] = result
	}

}

func i(slice []int) {
	for i := 0; i < len(slice); i++ {
		result := 1
		for j := 0; j < slice[i]; j++ {
			result = result * slice[i]
		}
		slice[i] = result
	}

}

func sumation(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	result := 0
	for i := 0; i < len(slice); i++ {
		result += slice[i]
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	power(nums, 3)

	fmt.Println(nums)

	nums2 := []int{1, 2, 3, 4, 5}

	i(nums2)

	fmt.Println(nums2)

	sum := sumation(nums2)

	fmt.Println(sum)
}
