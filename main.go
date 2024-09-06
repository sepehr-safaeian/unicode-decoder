package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root := "source"
}

func walkDir(root string) error {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".properties") {
			fmt.Println("Processing file:", path)
			if err := processFile(path); err != nil {
				return err
			}
		}
		return nil
	})
	return err
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

func decodeUnicode(s string) string {
	runes := []rune{}
	for len(s) > 0 {
		if strings.HasPrefix(s, `\u`) && len(s) >= 6 {
			var r rune
			fmt.Sscanf(s[2:6], "%04x", &r)
			runes = append(runes, r)
			s = s[6:]
		} else {
			runes = append(runes, rune(s[0]))
			s = s[1:]
		}
	}
	return string(runes)
}