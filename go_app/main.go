// Created by: Dominic M.
// Created on: April 2023
//
// This program finds the sum of natural numbers
package main

import (
	"fmt"
)

func main() {

	var number int
	var answer int
	var counter int

	// input
	fmt.Println("This program finds the sum of natural numbers")
	fmt.Println()
	fmt.Print("Enter a number: ")
	fmt.Println()
	fmt.Scanln(&number)

	// process
	for counter <= number {
		answer += counter
		counter++
	}

	fmt.Println()
	fmt.Print("The sum of all positive numbers from 0 to ", number, " is ", answer, ".")
	fmt.Println()
	fmt.Print("\nDone.")
}
