/*
 * File: linked_queue.go
 * Author: Franz Zbinden
 * Course: COTI 4039-LH1
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

type Container[T any] interface {
	IsEmpty() bool
	Contains(target T) bool
	All() iter.Seq[T]
}

// NewQueue returns an empty linked queue.
func NewLinkedQueue[T comparable]() *LinkedQueue[T] {
	return &LinkedQueue[T]{
		front: *node[T]
		rear:  *node[T]
	}
}

// IsEmpty determines whether the queue has zero items.
func (q *LinkedQueue[T]) IsEmpty() bool {
	return q.front == nil
}

// enqueue an item at the rear of the queue.
func (q *LinkedQueue[T]) Enqueue(item T) { // insertar dato al queue
	q.rear.next = &node[T]{data: item}
	q.rear = q.rear.next
}

// Pop removes the item at the top of the stack and returns the
// item or an error if the stack is empty.
func (q *LinkedQueue[T]) Dequeue() (item T, err error) {
	if q.IsEmpty() {
		var empty T
		return empty, errors.New("empty queue")
	}
	item = q.front.data
	q.front = q.front.next
	return item, nil
}

// Contains determines whether the stack contains the target item.
func (stk *LinkedQueue[T]) Contains(target T) bool {
	for curr := stk.top; curr != nil; curr = curr.next {
		if curr.data == target {
			return true
		}
	}
	return false
}

// All returns an iterator over the sequence of stack items.
func (stk *LinkedQueue[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for curr := stk.top; curr != nil; curr = curr.next {
			if !yield(curr.data) {
				return
			}
		}
	}
}
