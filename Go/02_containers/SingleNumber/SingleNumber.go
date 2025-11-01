/* File: my_slice_fns.go
 * Author: Franz Zbinden
 * Purpose: This program demonstrates how to define some slice functions.
 */

package main

import (
	"fmt"
)

func singleDigit(arr [5]int) int {
	for i := 0; i < len(arr); i++ {
		b := 0
		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] {
				b++
			}
		}
		if b == 1 {
			return arr[i]
		}
	}
	return -1
}

func main() {
	var numbers = [5]int{10, 20, 30, 10, 20}

	odd := singleDigit(numbers)

	fmt.Println(odd)
}
