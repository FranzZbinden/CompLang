/*
 * File: map_demo.go
 * Author: Franz Zbinden García
 * Course: COTI 4039-LH1
 * Date: 09/22/2025
 * Purpose: This program demonstrates how to define and use maps.
 */

package main

import "fmt"

func printMap(Map map[string]string) {
	fmt.Println("Spanish \t English")

	for spanish, english := range Map {
		fmt.Println(spanish, " \t", english)
	}

}

func main() {
	engToSpan := map[string]string{
		"Hola":    "Hello",
		"Manzana": "Apple",
		"Carro":   "Car",
	}

	fmt.Println(engToSpan)

	fmt.Println()

	fmt.Println(engToSpan["Hola"])

	fmt.Println()

	printMap(engToSpan)

}
