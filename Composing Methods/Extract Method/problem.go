package main

import "fmt"

// You have a code fragment that can be grouped together.

// Why Refactor:
//The more lines found in a method, the harder it’s to figure out what the method does. This is the main reason for this refactoring.

func printOwing() {
	printBanner()

	// Print details.
	fmt.Println("name: ")
	fmt.Println("amount: ", getOutstanding())
}

func getOutstanding() float64 {
	return 3.1416
}
func printBanner() {
	// Priting banner
}
