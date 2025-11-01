/*
 * File: queue.go
 * Author: Franz Zbinden
 * Date: 10/20/2025
 * Purpose: This is the interface for a generic queue.
 */

package container

// Queue represents a container where addition and removal occurs at opposite ends.
type Queue[T comparable] interface {
	Container[T]
	Enqueue(item T)
	Dequeue() (item T, err error)
}
