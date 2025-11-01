package container

/*
 * File: stack_test.go
 * Author: Franz Zbinden
 * Date: 10/20/2025
 * Purpose: This is the set of unit tests for generic stacks.
 * Note: Use the 'go test -v' command to run all tests.
 */

import "testing"

// Tests that a new stack is empty.
func TestNewStack(t *testing.T) {
	var stk Stack[int] = NewLinkedStack[int]()

	if !stk.IsEmpty() {
		t.Error("Expected new stack to be empty")
	}
}

// Tests that pushing an element results in a non-empty stack.
func TestPushElement(t *testing.T) {
	var stk Stack[int] = NewLinkedStack[int]()
	stk.Push(10)

	if stk.IsEmpty() {
		t.Error("Expected stack to be non-empty after push")
	}
}

// Tests that a non-empty stack contains only its elements.
func TestContainsElements(t *testing.T) {
	var stk Stack[int] = NewLinkedStack[int]()
	stk.Push(10)
	stk.Push(20)

	if !stk.Contains(10) {
		t.Errorf("Expected stack to contain 10")
	}

	if !stk.Contains(20) {
		t.Errorf("Expected stack to contain 20")
	}

	if stk.Contains(30) {
		t.Errorf("Expected stack to not contain 30")
	}
}

// Tests that a non-empty stack has its elements in reversed order.
func TestAllElements(t *testing.T) {
	var stk Stack[int] = NewLinkedStack[int]()
	stk.Push(10)
	stk.Push(20)
	stk.Push(30)

	expected := []int{30, 20, 10}
	index := 0
	for item := range stk.All() {
		if item != expected[index] {
			t.Errorf("Expected %d, got %d", expected[index], item)
		}
		index++
	}
	if index != len(expected) {
		t.Errorf("Expected %d items, got %d", len(expected), index)
	}
}

// Tests that elements are popped in reversed order from a non-empty stack.
func TestPopElements(t *testing.T) {
	var stk Stack[int] = NewLinkedStack[int]()
	stk.Push(10)
	stk.Push(20)

	item, err := stk.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if item != 20 {
		t.Errorf("Expected 20 but got %d", item)
	}

	item, err = stk.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if item != 10 {
		t.Errorf("Expected 10 but got %d", item)
	}

	_, err = stk.Pop()
	if err == nil {
		t.Errorf("Expected error when popping from empty stack")
	}
}
