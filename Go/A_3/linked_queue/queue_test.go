/*
 * File: queue_test.go
 * Author: Franz Zbinden
 * Date: 10/20/2025
 * Purpose: This is the set of unit tests for generic queues.
 * Note: Use the 'go test -v' command to run all tests.
 */

package container

import "testing"

// Declares a new queue for testing.
var que Queue[int]

// Performs common setup for every test.
func setup() {
	que = NewLinkedQueue[int]()
}

// Tests that a new queue is empty.
func TestNewQueue(t *testing.T) {
	setup()

	if !que.IsEmpty() {
		t.Error("Expected new queue to be empty")
	}
}

// Tests that adding an element results in a non-empty queue.
func TestEnqueueElement(t *testing.T) {
	setup()
	que.Enqueue(10)

	if que.IsEmpty() {
		t.Error("Expected queue to be non-empty after enqueue")
	}
}

// Tests that a non-empty queue contains only its elements.
func TestContainsElements(t *testing.T) {
	setup()
	que.Enqueue(10)
	que.Enqueue(20)

	if !que.Contains(10) {
		t.Error("Expected queue to contain 10")
	}

	if !que.Contains(20) {
		t.Error("Expected queue to contain 20")
	}

	if que.Contains(30) {
		t.Error("Expected queue to not contain 30")
	}
}

// Tests that a non-empty queue has its elements in insertion order.
func TestAllElements(t *testing.T) {
	setup()
	que.Enqueue(10)
	que.Enqueue(20)
	que.Enqueue(30)

	expected := []int{10, 20, 30}
	index := 0
	for item := range que.All() {
		if item != expected[index] {
			t.Errorf("Expected %d, got %d", expected[index], item)
		}
		index++
	}
	if index != len(expected) {
		t.Errorf("Expected %d items, got %d", len(expected), index)
	}
}

// Tests that elements are removed in insertion order from a non-empty queue.
func TestDequeueElements(t *testing.T) {
	setup()
	que.Enqueue(10)
	que.Enqueue(20)

	item, err := que.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if item != 10 {
		t.Errorf("Expected 10 but got %d", item)
	}

	item, err = que.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if item != 20 {
		t.Errorf("Expected 20 but got %d", item)
	}

	_, err = que.Dequeue()
	if err == nil {
		t.Error("Expected error when removing from an empty queue")
	}
}
