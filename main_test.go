package main

import (
	"os"
	"testing"
)

func TestAnalyzeFile(t *testing.T) {
	// Create a temporary file for testing
	file, err := os.CreateTemp("", "test_file.txt")
	if err != nil {
		t.Fatal("Failed to create temporary file:", err)
	}
	defer os.Remove(file.Name())

	// Write test content to the file
	content := "Hello, world!\nThis is a test file for analysis.\n"
	_, err = file.WriteString(content)
	if err != nil {
		t.Fatal("Failed to write to temporary file:", err)
	}
	file.Close()

	// Open the test file
	file, err = os.Open(file.Name())
	if err != nil {
		t.Fatal("Failed to open temporary file:", err)
	}
	defer file.Close()

	// Analyze the test file
	wordCount, lineCount, charCount := analyzeFile(file)

	// Check the results
	expectedWordCount := 8
	expectedLineCount := 2
	expectedCharCount := len(content) - 1 // Excluding the newline character

	if wordCount != expectedWordCount {
		t.Errorf("Word count mismatch. Expected: %d, got: %d", expectedWordCount, wordCount)
	}

	if lineCount != expectedLineCount {
		t.Errorf("Line count mismatch. Expected: %d, got: %d", expectedLineCount, lineCount)
	}

	if charCount != expectedCharCount {
		t.Errorf("Character count mismatch. Expected: %d, got: %d", expectedCharCount, charCount)
	}
}
