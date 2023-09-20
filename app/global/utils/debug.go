package utils

import "fmt"

// Void takes in everything and sends everthing into the void
func Void(a ...interface{}) {} // What for? Simply to satisfy compiler and linter when they complain about unused variables...

// SuperLog prints everything to the stdout with two new-liners on two ends for better readability
func SuperLog(a ...interface{}) {
	fmt.Println("\n\n ", a, "\n\n ")
}
