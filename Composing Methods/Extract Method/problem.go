package main

import "fmt"

//You have a code fragment that can be grouped together.

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
