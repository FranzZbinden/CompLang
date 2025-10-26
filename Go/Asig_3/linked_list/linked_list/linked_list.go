/*
 * File: linked_list.go
 * Author: Franz Zbinden García
 * Course: COTI 4039-VH1
 * Date: 10/19/2025
 * Purpose: This program demonstrates how to define and use a generic
 *          linked list & assigment.
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

// a.	A method that returns the 0-based position of the given value in a generic list:
func (lst *list[T]) positionOf(target T) int {
	current := 0

	for curr := lst; curr != nil; curr = curr.next {
		if curr.value == target {
			return current
		}
		current += 1
	}
	return -1
}

// b.	A method that appends a value at the end of a generic list:
func (lst *list[T]) append(value T) *list[T] {
	if lst.next != nil {
		return lst.next.append(value)
	}
	lst.next = &list[T]{next: nil, value: value}
	return lst
}

// c.	A function that returns the returns the minimum and maximum of a list of numbers:
func extrema[N number](lst *list[N]) (min, max N) {
	max_temp := lst.value
	min_temp := lst.value

	for curr := lst; curr != nil; curr = curr.next {
		if curr.value > max_temp {
			max_temp = curr.value
		}
		if curr.value < min_temp {
			min_temp = curr.value
		}
	}
	return min_temp, max_temp
}

// Prepends the given value to the given list. (adds a node to the list)
func cons[T comparable](value T, lst *list[T]) *list[T] {
	return &list[T]{value, lst}
}

// Determines whether the list is empty.
func (lst *list[T]) isEmpty() bool {
	return lst == nil // If the pointer is nil, not pointing at anything then is empty
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

	target := 20
	fmt.Println("\nThe position of", target, "is", numbers.positionOf(target))

	// Using append to add 75
	add_num := 75
	numbers.append(add_num)

	fmt.Println("After appending", add_num, "the list of numbers is", numbers)

	// func extrema[N number](lst *list[N]) (min, max N) {
	min, max := extrema(numbers)
	fmt.Println("The extrema are ", min, "and", max)

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
