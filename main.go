package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define a flag for the -c option
	countBytes := flag.Bool("c", false, "count bytes")
	flag.Parse()

	// Check if the countBytes flag is provided
	if !*countBytes {
		fmt.Println("Usage: ccwc -c <filename>")
		return
	}

	// Get the file name from the arguments
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Usage: ccwc -c <filename>")
		return
	}
	fileName := args[0]

	// Read the file
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", fileName, err)
		return
	}

	// Get the number of bytes
	numBytes := len(fileData)

	// Print the result
	fmt.Printf("%d %s\n", numBytes, fileName)
}