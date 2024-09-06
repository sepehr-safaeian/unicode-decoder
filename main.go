package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Main function to initiate the directory traversal and file processing.
func main() {
	rootDir := "source" // Define the root directory to start processing.
	if err := walkDirectory(rootDir); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Processing completed successfully.")
	}
}

// walkDirectory traverses through all the files in the given directory and processes .properties files.
func walkDirectory(root string) error {
	// Walk through the directory tree starting from the root path.
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // If there's an error while accessing the file/directory, return it.
		}

		// Process only files with .properties extension.
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".properties") {
			fmt.Printf("Processing file: %s\n", path)
			if err := processFile(path); err != nil {
				return err // If file processing fails, return the error.
			}
		}
		return nil
	})
}

// processFile reads a file, decodes any Unicode escape sequences, and writes the decoded content back.
func processFile(filePath string) error {
	// Open the file for reading.
	file, err := os.Open(filePath)
	if err != nil {
		return err // Return error if file can't be opened.
	}
	defer file.Close() // Ensure the file is closed when the function finishes.

	// Create a buffered reader to read the file line by line.
	reader := bufio.NewReader(file)
	var decodedLines []string
	// Regular expression to match Unicode escape sequences in the form of \uXXXX.
	unicodeRegex := regexp.MustCompile(`\\u[0-9A-Fa-f]{4}`)

	for {
		// Read each line from the file.
		line, err := reader.ReadString('\n')
		if err != nil && err.Error() != "EOF" {
			return err // Return any error except EOF.
		}
		// Trim newline and carriage return characters from the end of the line.
		line = strings.TrimRight(line, "\n\r")

		// Replace all Unicode escape sequences with their decoded character equivalents.
		decodedLine := unicodeRegex.ReplaceAllStringFunc(line, decodeUnicode)
		decodedLines = append(decodedLines, decodedLine)

		if err != nil { // EOF case, break after processing the last line.
			break
		}
	}

	// Write the decoded lines back to the original file.
	return writeToFile(filePath, decodedLines)
}

// writeToFile writes the processed lines to the original file.
func writeToFile(filePath string, lines []string) error {
	// Create (or truncate) the file for writing the output.
	outputFile, err := os.Create(filePath)
	if err != nil {
		return err // Return error if the file can't be created or opened.
	}
	defer outputFile.Close() // Ensure the file is closed after writing.

	// Write each line to the output file.
	for _, line := range lines {
		_, err := outputFile.WriteString(line + "\n")
		if err != nil {
			return err // Return error if writing to file fails.
		}
	}

	return nil
}

// decodeUnicode takes a Unicode escape sequence in the form of \uXXXX and converts it to the actual character.
func decodeUnicode(unicodeSeq string) string {
	var runes []rune
	// While there are characters left in the string, decode them.
	for len(unicodeSeq) > 0 {
		// If the string starts with \u and has at least 6 characters, it's a Unicode sequence.
		if strings.HasPrefix(unicodeSeq, `\u`) && len(unicodeSeq) >= 6 {
			var r rune
			// Convert the hexadecimal Unicode sequence (XXXX) to a rune.
			fmt.Sscanf(unicodeSeq[2:6], "%04x", &r)
			runes = append(runes, r)    // Add the decoded rune to the runes slice.
			unicodeSeq = unicodeSeq[6:] // Move to the next part of the string.
		} else {
			// If it's not a Unicode sequence, just append the character.
			runes = append(runes, rune(unicodeSeq[0]))
			unicodeSeq = unicodeSeq[1:] // Move to the next character.
		}
	}
	// Return the decoded string.
	return string(runes)
}
