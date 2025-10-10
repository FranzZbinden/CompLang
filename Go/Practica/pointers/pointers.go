/*
 * File: payroll.go
 * Author: Franz Zbinden Garcia
 * Course: COTI 4039-VH1
 * Date: 09/30/2025
 * Purpose: Part 3 of the Assigment.
 */

package main

import "fmt"

func sum(p *int) {
	*p = *p + *p

}

func main() {
	a := 1
	b := 2

	fmt.Println(a, b)

	fmt.Println()

	fmt.Println(&a, &b)

	pointerA := &a
	pointerB := &b

	fmt.Println()

	fmt.Println(pointerA, pointerB)

	fmt.Println()

	fmt.Println(*pointerA, *pointerB)

	p := 2

	sum(&p)

	fmt.Println(p)

}
