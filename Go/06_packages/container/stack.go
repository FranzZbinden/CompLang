/*
 * File: stack.go
 * Author: Franz Zbinden
 * Date: 10/20/2025
 * Purpose: This is the interface for a generic stack.
 */

package container

// Stack represents a container where addition and removal occurs at the same end.
type Stack[T comparable] interface {
	Container[T]
	Push(item T)
	Pop() (item T, err error)
}
