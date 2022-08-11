package main

import "fmt"

func printOwingSol() {
	printBanner()

	// Print details.
	printDetails(getOutstanding())
}

func printDetails(outstanding float64) {
	fmt.Println("name: ")
	fmt.Println("amount: ", outstanding)
}
