package container

/*
 * File: container.go
 * Author: Franz Zbinden
 * Date: 10/20/2025
 * Purpose: This is the interface for a generic container and the error for
 *          an empty container.
 */

// Package container has interfaces and implementations of generic containers.

import (
	"iter"
)

// Container is an interface for a generic group of items.
type Container[T comparable] interface {
	IsEmpty() bool
	Contains(target T) bool
	All() iter.Seq[T]
}

// EmptyError represents the error for an empty container.
type EmptyError struct {
	message string
}

// Returns the message string associated with the error.
func (e *EmptyError) Error() string {
	return e.message
}
