// Copyright (c) 2021 Daniel Pawelko All rights reserved
//
// Created by: Daniel Pawelko
// Created on: May 2021
// This file contains GO program that returns if you should go to the movies

package main

import (
	"fmt"
)

// This main function will ask user for age and return what the user can watch
func main() {
	// Defining variables
	var age int
	var day string

	fmt.Println("Should you go to the movies?")
	fmt.Println()

	// User Input
	fmt.Print("Age: ")
	fmt.Scanln(&age)
	fmt.Println()

	fmt.Print("Day of the week: ")
	fmt.Scanln(&day)
	fmt.Println()

	// Return what user can watch
	if (day == "tuesday" || day == "thursday") || (age > 12 && age < 21) {
		fmt.Println("You should go to the museum because you will get a discount")
	} else {
		fmt.Println("You should not go to the museum today because you will not get a discount")
	}
}
