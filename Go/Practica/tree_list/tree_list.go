/*
 * File: binary_tree.go
 * Author: Franz Zbinden
 * Course: COTI 4039-LH1
 * Date: 10/15/2025
 * Purpose: This program demonstrates how to define and use a generic
 *          binary search tree.
 */
// bst al principio esta apuntando al root
package main

import (
	"cmp"
	"fmt"
	"iter"
	"strings"
)

// Represents a tree with a value and pointers to the left and right subtrees.
type tree[T cmp.Ordered] struct {
	left  *tree[T]
	value T
	right *tree[T]
}

type tree2[T cmp.Ordered] struct {
	left  *tree2[T]
	value T
	right *tree2[T]
}

func newTree2[T cmp.Ordered]() *tree[T] {
	return nil
}

// funcion
// Creates a new empty tree.
func newTree[T cmp.Ordered]() *tree[T] {
	return nil
}

// metodo
// Determines whether the tree is empty.
func (bst *tree[T]) isEmpty() bool {
	return bst == nil
}

func (bst *tree[T]) isEmpty2() bool {
	return bst == nil
}

// Returns the number of elements in the tree.
func (bst *tree[T]) size() int {
	if bst == nil {
		return 0
	}
	return 1 + bst.left.size() + bst.right.size()
}

func (bst *tree[T]) size2() int {
	if bst == nil {
		return 0
	}
	return 1 + bst.left.size() + bst.right.size()
}

// Sums the elements in the tree.
func (bst *tree[T]) sum() T {
	if bst == nil {
		var zero T // en ves de devolver un 0 devuelve una variable generica
		return zero
	}
	return bst.value + bst.left.sum() + bst.right.sum()
}

func (bst *tree[T]) sum2() T {
	if bst == nil {
		var zero T
		return zero
	}
	return bst.value + bst.left.sum2() + bst.right.sum2()
}

// Inserts an element to the tree.
func (bst *tree[T]) insert(elem T) *tree[T] {
	if bst == nil { // este no se podria poner fuera de los if statments
		return &tree[T]{left: nil, value: elem, right: nil}
	}
	if elem < bst.value {
		bst.left = bst.left.insert(elem) //devolvi el arbol cambiado por la izquierda
	} else if elem > bst.value {
		bst.right = bst.right.insert(elem) //devolvi el arbol cambiado por la derecha
	}
	return bst

}

func (bst *tree[T]) insert2(elem T) *tree[T] {
	if bst == nil {
		return &tree[T]{value: elem}
	}
	if bst.value > elem {
		bst.left = bst.left.insert(elem)
	} else if bst.value < elem {
		bst.right = bst.right.insert(elem)
	}
	return bst

}

// Determines whether the tree contains the given target.
func (bst *tree[T]) contains(target T) bool {
	if bst == nil { // este no se podria poner fuera de los if statments
		return false
	}
	if target < bst.value {
		return bst.left.contains(target)
	} else if target > bst.value {
		return bst.right.contains(target)
	}
	return true
}

// Returns the string representation of the tree.
func (bst *tree[T]) String() string {
	sb := strings.Builder{}
	sb.WriteString("tree{ ") //escribes tree{ al principio del string

	bst.putIntoString(&sb)

	sb.WriteString("}") //escribes } al final del string
	return sb.String()
}

// Puts the tree elements into the given string builder.
func (bst *tree[T]) putIntoString(sb *strings.Builder) {
	if bst == nil {
		return
	}
	bst.left.putIntoString(sb)
	sb.WriteString(fmt.Sprintf("%v ", bst.value))
	bst.right.putIntoString(sb)
}

// Returns an iterator over a sequence of all elements in the list.
func (bst *tree[T]) all() iter.Seq[T] {
	return func(yield func(T) bool) {
		bst.putIntoSeq(yield)
	}
}

// Puts the tree elements into a sequence using the given yield function.
func (bst *tree[T]) putIntoSeq(yield func(T) bool) bool {
	return bst == nil ||
		bst.left.putIntoSeq(yield) &&
			yield(bst.value) &&
			bst.right.putIntoSeq(yield)

}

// Represents a list with a value and a pointer to the next sublist.
type list[T comparable] struct {
	value T
	next  *list[T]
}

// Prepends the given value to the given list. (adds a node to the list)
func cons[T comparable](value T, lst *list[T]) *list[T] {
	return &list[T]{value, lst}
}

func cons2[T comparable](val T, lst *list[T]) *list[T] {
	return &list[T]{value: val, next: lst}
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

func TreeToLinkedList[T cmp.Ordered](root *tree[T]) *list[T] {
	if root == nil {
		return nil
	}

	// We'll build the list in reverse order, then return it.
	var head *list[T]

	// Helper function: traverses the tree in order and adds nodes to the list.
	var build func(*tree[T])
	build = func(node *tree[T]) {
		if node == nil {
			return
		}
		// Traverse right first so that we can prepend and still get left-to-right order
		build(node.right)
		head = cons(node.value, head)
		build(node.left)
	}

	build(root)
	return head
}

func TreeToSlice[T cmp.Ordered](root *tree[T]) []T {
	if root == nil {
		return nil
	}
	// Traverse left, then node, then right (in-order)
	left := TreeToSlice(root.left)
	right := TreeToSlice(root.right)
	return append(append(left, root.value), right...)
}

//menor, mayor funciones, calcular la profundidad del arbol

// creas un apuntador a un arbol para
// Starts the execution of the program.
func main() {
	numbers := newTree[int](). // los parentesis son porque se esta llamando
					insert(30).
					insert(10).
					insert(50).
					insert(40).
					insert(20)

	fmt.Println("The tree of numbers is", numbers)
	fmt.Println("\nThe root value is", numbers.value)
	fmt.Println("The left subtree is", numbers.left)
	fmt.Println("The right subtree is", numbers.right)

	fmt.Println("\nIts size is", numbers.size())
	fmt.Println("Is it empty?", numbers.isEmpty())
	fmt.Println("The sum of its elements is", numbers.sum())
	fmt.Println("Does it contain 20?", numbers.contains(20))

	fmt.Println("\nThe elements, one per line...")
	for elem := range numbers.all() {
		fmt.Println(elem)
	}
}
