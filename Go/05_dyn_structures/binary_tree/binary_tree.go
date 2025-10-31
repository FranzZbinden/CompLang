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
	if elem > bst.value {
		bst.right = bst.right.insert2(elem)
	} else if elem < bst.value {
		bst.left = bst.left.insert2(elem)
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

func (bst *tree[T]) contains2(target T) bool {
	if bst == nil {
		return false
	}
	if target > bst.value {
		return bst.right.contains(target)
	} else if target < bst.value {
		return bst.left.contains(target)
	}
	return true
}

// Returns the string representation of the tree.
func (bst *tree[T]) String() string {
	sb := strings.Builder{}
	sb.WriteString("tree{ ")

	bst.putIntoString(&sb)

	sb.WriteString("}")
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

// Collects all elements of the tree into a slice (in-order traversal).
func (bst *tree[T]) putIntoSlice2() []T {
	if bst == nil {
		return nil
	}
	left := bst.left.putIntoSlice2()   // type pointer to tree
	right := bst.right.putIntoSlice2() // type pointer to tree

	result := append(left, bst.value)
	result = append(result, right...)
	return result
}

func (bst *tree[T]) slice_to_tree(arr []T) *tree[T] {
	if len(arr) == 0 {
		return bst
	}
	val := arr[0]
	if bst == nil {
		bst = &tree[T]{value: val}
	} else if val > bst.value {
		bst.right = bst.right.slice_to_tree(arr[1:])
	} else {
		bst.left = bst.left.slice_to_tree(arr[1:])
	}
	return bst.slice_to_tree(arr[1:])
}

// Starts the execution of the program.
func main() {
	numbers := newTree[int]().
		insert(30).
		insert(10).
		insert(50).
		insert(40).
		insert(20)

	arr := []int{1, 2, 3, 4}
	numbers2 := newTree[int]()
	numbers2.slice_to_tree(arr)
	fmt.Println(numbers2)

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
