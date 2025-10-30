/*
 * File: slice_demo.go
 * Author: Franz Zbinden
 * Course: COTI 4039-LH1
 * Date: 09/17/2025
 * Purpose: This program demonstrates how to define and use slices.
 */

package main

import (
	"fmt"
	"slices"
)

// Sums the elements of a slice.
func sum(slice []int) int {
	total := 0
	for _, elem := range slice {
		total += elem
	}
	return total
}

// Starts the execution of the program.
func main() {
	numbers := []int{30, 10, 50, 40, 20}

	fmt.Println("The slice of numbers is", numbers)

	fmt.Println("\nThese are the elements, one per line: ")
	for index, number := range numbers {
		fmt.Println(index, ":", number)
	}

	fmt.Println("\nIts length is", len(numbers))
	fmt.Println("Its capacity is", cap(numbers))
	fmt.Println("The sum of its elements is", sum(numbers))

	fmt.Println("\nIts minimum element is", slices.Min(numbers))
	fmt.Println("Its maximum element is", slices.Max(numbers))

	fmt.Println("\nDoes it contain 20?", slices.Contains(numbers, 20))
	fmt.Println("The index of 20 is", slices.Index(numbers, 20))
	fmt.Println("The element at index #4 is", numbers[4])

	fmt.Println("\nIs it sorted?", slices.IsSorted(numbers))
	slices.Sort(numbers)
	fmt.Println("The sorted slice is", numbers)
	fmt.Println("Is it sorted?", slices.IsSorted(numbers))

	index, found := slices.BinarySearch(numbers, 20)
	fmt.Println("\nDoes it contain 20?", found)
	fmt.Println("The index of 20 is", index)
	fmt.Println("The element at index #1 is", numbers[1])

	index, found = slices.BinarySearch(numbers, 25)
	fmt.Println("\nDoes it contain 25?", found)
	fmt.Println("Index where 25 should be placed?", index)

	slices.Reverse(numbers)
	fmt.Println("\nThe reversed slice is", numbers)

	numbers2 := slices.Clone(numbers)
	fmt.Println("The slice clone is", numbers2)

	clear(numbers2)
	fmt.Println("\nAfter clearing, the slice clone is", numbers2)
	fmt.Println("The original slice of numbers still is", numbers)

	evenSquares := make([]int, 0)
	for num := 1; num <= 20; num++ {
		if num%2 == 0 {
			evenSquares = append(evenSquares, num*num)
		}
	}
	fmt.Println("\nThe first ten even squares are", evenSquares)
	fmt.Println("Its four middle elements are", evenSquares[3:7])

	deleted := slices.Delete(evenSquares, 3, 7)
	fmt.Println("\nAfter deleting its four middle elements, the new slice is", deleted)

	inserted := slices.Insert(deleted, 2, 10, 20, 30)
	fmt.Println("After inserting some elements, the new slice is", inserted)
}
