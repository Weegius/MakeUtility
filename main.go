package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Prompt the user for the file path
	fmt.Println("Text File Analyzer")
	fmt.Println("-------------------------------")
	fmt.Println("Please enter the path to the text file:")

	var filePath string
	fmt.Scanln(&filePath)

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file, please try again.")
		log.Fatal(err)
	}
	defer file.Close()

	// Analyze the file and print the statistics
	wordCount, lineCount, charCount := analyzeFile(file)

	fmt.Println("Analyzing the file...")
	fmt.Printf("Word Count: %d\n", wordCount)
	fmt.Printf("Line Count: %d\n", lineCount)
	fmt.Printf("Character Count: %d\n", charCount)
}

func analyzeFile(file *os.File) (int, int, int) {
	wordCount := 0
	lineCount := 0
	charCount := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Count the number of lines
		lineCount++

		// Count the number of characters
		charCount += len(line)

		// Split the line into words
		words := strings.Fields(line)
		wordCount += len(words)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return wordCount, lineCount, charCount
}
