/*
 * File: postfix_eval_io.go
 * Author: Franz Zbinden
 * Date: 10/20/2025
 * Purpose: This program evaluates a postfix expression.
 */

// Package main contains the main program of the postfix evaluator.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"uprb.edu/container"
)

// ExpressionError represents the error when evaluating an expression.
type ExpressionError struct {
	message string
}

// Returns the message string associated with the error.
func (e *ExpressionError) Error() string {
	return e.message
}

// Function isOperator determines whether the given token is an arithmetic
// operator.
func isOperator(token string) bool {
	switch token {
	case "+", "-", "*", "/":
		return true
	}
	return false
}

// Function evalBinaryExpr returns the result of evaluating a binary
// expression or an error for an unknown operator.
func evalBinaryExpr(operand1, operand2 int, operator string) (int, error) {
	switch operator {
	case "+":
		return operand1 + operand2, nil
	case "-":
		return operand1 - operand2, nil
	case "*":
		return operand1 * operand2, nil
	case "/":
		return operand1 / operand2, nil
	}
	return 0, &ExpressionError{operator + " is an invalid operator"}
}

// Function evalPostfixExpr returns the result of evaluating a postfix
// expression or an error for an invalid expression.
func evalPostfixExpr(postfix string) (int, error) {
	operands := container.NewLinkedStack[int]()
	tokens := strings.SplitSeq(postfix, " ")

	for token := range tokens {
		if isOperator(token) {
			operand2, err := operands.Pop()
			if err != nil {
				return 0, &ExpressionError{"missing second operand"}
			}

			operand1, err := operands.Pop()
			if err != nil {
				return 0, &ExpressionError{"missing first operand"}
			}

			result, err := evalBinaryExpr(operand1, operand2, token)
			if err != nil {
				return 0, err
			}

			operands.Push(result)
		} else {
			operand, err := strconv.Atoi(token)
			if err != nil {
				return 0, &ExpressionError{token + " is an invalid token"}
			}

			operands.Push(operand)
		}
	}
	result, err := operands.Pop()
	if err != nil {
		return 0, &ExpressionError{"no result was found"}
	}

	if !operands.IsEmpty() {
		return 0, &ExpressionError{"missing operator"}
	}
	return result, nil
}

// Function main is the entry point of this program.
func main() {
	fmt.Print("Enter a postfix expression: ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		result, err := evalPostfixExpr(scanner.Text())
		if err != nil {
			fmt.Println("Error: " + err.Error())
			os.Exit(1) // failure
		}

		fmt.Println("\nThe expression equals", result)
	}
}
