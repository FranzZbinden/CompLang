/*
 * File: linked_list.go
 * Author: Franz Zbinden
 * Date: 10/08/2025
 * Purpose: This program demonstrates how to define and use a generic
 *          linked list.
 */

package main

import (
	"fmt"
	"iter"
	"strings"
)

// Represents a list with a value and a pointer to the next sublist.
type list[T comparable] struct {
	value T
	next  *list[T]
}

// Prepends the given value to the given list. (adds a node to the list)
func cons[T comparable](value T, lst *list[T]) *list[T] {
	return &list[T]{value, lst}
}

// Determines whether the list is empty.
func (lst *list[T]) isEmpty() bool {
	return lst == nil
}

// Returns the number of elements in a list.
func (lst *list[T]) length() int {
	len := 0
	for curr := lst; curr != nil; curr = curr.next {
		len++
	}
	return len
}

// Determines whether the list contains the given target.
func (lst *list[T]) contains(target T) bool {
	for curr := lst; curr != nil; curr = curr.next {
		if target == curr.value {
			return true
		}
	}
	return false
}

func (lst *list[T]) contains2(target T) bool {
	if lst == nil {
		return false
	}
	if lst.value == target {
		return true
	}
	return lst.next.contains2(target)
}

// Returns the string representation of the list.
func (lst *list[T]) String() string {
	sb := strings.Builder{}
	sb.WriteString("list{ ")
	for curr := lst; curr != nil; curr = curr.next {
		sb.WriteString(fmt.Sprintf("%v ", curr.value))
	}
	sb.WriteString("}")
	return sb.String()
}

// Returns an iterator over a sequence of all elements in the list.
func (lst *list[T]) all() iter.Seq[T] {
	return func(yield func(T) bool) {
		for curr := lst; curr != nil; curr = curr.next {
			if !yield(curr.value) {
				return
			}
		}
	}
}

// Collects all elements of the tree into a slice (in-order traversal).
func (lst *list[T]) putIntoSlice() []T {
	if lst == nil {
		return nil
	}
	return append([]T{lst.value}, lst.next.putIntoSlice()...)
}

// Starts the execution of the program.
func main() {
	numbers := cons(30, cons(10, cons(50, cons(40, cons(20, nil)))))

	fmt.Println("The list of numbers is", numbers)
	fmt.Println("\nThe first value is", numbers.value)
	fmt.Println("The next sublist is", numbers.next)

	fmt.Println("\nIts length is", numbers.length())
	fmt.Println("Is it empty?", numbers.isEmpty())
	fmt.Println("Does it contain 20?", numbers.contains(20))

	fmt.Println("\nThese are the elements, one per line...")
	for elem := range numbers.all() {
		fmt.Println(elem)
	}

	fmt.Println("\nThe sum of its elements is", sum(numbers))

}

// Represents the union of common number types.
type number = interface {
	~int | ~float32 | ~float64
}

// Sums the elements in the given list of numbers.
func sum[N number](lst *list[N]) N {
	var total N
	for curr := lst; curr != nil; curr = curr.next {
		total += curr.value
	}
	return total
}

func sumation[N number](lst *list[N]) N {
	var num N
	for curr := lst; curr != nil; curr = curr.next {
		num += curr.value
	}
	return num
}
