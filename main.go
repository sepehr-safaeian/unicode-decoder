package main

import (
	"fmt"
	"os"
)

func main() {
	root := "source"
}

func processFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	vat outputLines []string
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\\u[0-9A-Fa-f]{4}`)
	for scanner.Scan() {
		line := scanner.Text()
		decodedLine := re.ReplaceAllStringFunc(line, decodeUnicode)
		outputLines = append(outputLines, decodedLine)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	outputFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	for _, line := range outputLines {
		_, err := outputFile.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
