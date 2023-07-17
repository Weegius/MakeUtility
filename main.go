package main

import (
	"fmt"
	"log"
	"os"
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
		log.Fatal(err)
	}
	defer file.Close()

	// Analyze the file
	wordCount := 0
	lineCount := 0
	charCount := 0

	fmt.Println("Analyzing the file...")

	// Next Implement basic analysis of the file

	// Print the statistics
	fmt.Printf("Word Count: %d\n", wordCount)
	fmt.Printf("Line Count: %d\n", lineCount)
	fmt.Printf("Character Count: %d\n", charCount)
}
