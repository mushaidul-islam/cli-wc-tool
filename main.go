package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	// Define the flags
	countBytes := flag.Bool("b", false, "count bytes")
	countLines := flag.Bool("l", false, "count lines")
	countWords := flag.Bool("w", false, "count words")
	countChars := flag.Bool("c", false, "count characters")
	flag.Parse()

	defaultFlag := !*countBytes && !*countLines && !*countWords && !*countChars

	// Get the file name from the arguments
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Usage: ccwc -c <filename>")
		return
	}
	fileName := args[0]

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", fileName, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var bytesCount, lineCount, wordCount, charCount int

	getLineCount := defaultFlag || *countLines
	getBytesCount := defaultFlag || *countBytes
	getWordsCount := defaultFlag || *countWords
	getCharsCount := defaultFlag || *countChars

	for scanner.Scan() {
		line := scanner.Text()

		if getLineCount {
			lineCount++
		}

		if getBytesCount {
			bytesCount += len(line) + 1 // To handle new-line character
		}

		if getWordsCount {
			wordCount += len(strings.Fields(line))
		}

		if getCharsCount {
			charCount += utf8.RuneCountInString(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	var result []string

	if getBytesCount {
		result = append(result, fmt.Sprintf("bytes: %d", bytesCount))
	}
	if getLineCount {
		result = append(result, fmt.Sprintf("lines: %d", lineCount))
	}
	if getWordsCount {
		result = append(result, fmt.Sprintf("words: %d", wordCount))
	}
	if getCharsCount {
		result = append(result, fmt.Sprintf("characters: %d", charCount))
	}

	fmt.Print(strings.Join(result, ", "))
}
