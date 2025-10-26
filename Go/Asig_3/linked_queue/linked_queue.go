/*
 * File: linked_queue.go
 * Author: Franz Zbinden
 * Course: COTI 4039-VH1
 * Date: 10/23/2025
 * Purpose: This is the implementation for a generic queue using links.
 */

package container

import (
	"errors"
	"iter"
)

// Type node represents a generic linked-queue node with its data and a
// pointer to the next node in the queue.
type node[T any] struct {
	data T
	next *node[T]
}

// LinkedStack represents a generic queue using links.
type LinkedQueue[T comparable] struct {
	front *node[T]
	rear  *node[T]
}

// container interface
type Container[T any] interface {
	IsEmpty() bool
	Contains(target T) bool
	All() iter.Seq[T]
}

// NewQueue returns an empty linked queue.
func NewLinkedQueue[T comparable]() *LinkedQueue[T] {
	return &LinkedQueue[T]{}
}

// IsEmpty determines whether the queue has zero items.
func (q *LinkedQueue[T]) IsEmpty() bool {
	return q.front == nil
}

// enqueue an item at the rear of the queue.
func (q *LinkedQueue[T]) Enqueue(item T) { // insertar dato al queue
	new_node := &node[T]{data: item}
	if q.IsEmpty() {
		q.front = new_node // nodo con nuevo item
		q.rear = new_node
	} else {
		q.rear.next = &node[T]{data: item}
		q.rear = q.rear.next
	}

}

// dequeue removes the item at the top of the queue and returns the
// item or an error if the queue is empty.
func (q *LinkedQueue[T]) Dequeue() (item T, err error) {
	if q.IsEmpty() {
		var empty T
		return empty, errors.New("empty queue")
	}
	item = q.front.data
	q.front = q.front.next
	return item, nil
}

// Contains determines whether the queue contains the target item.
func (q *LinkedQueue[T]) Contains(target T) bool {
	for curr := q.front; curr != nil; curr = curr.next {
		if curr.data == target {
			return true
		}
	}
	return false
}

// All returns an iterator over the sequence of queue items.
func (q *LinkedQueue[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for curr := q.front; curr != nil; curr = curr.next {
			if !yield(curr.data) {
				return
			}
		}
	}
}
