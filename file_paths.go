package main

import (
	"fmt"
	"os"
)

// processFilePaths takes as argument different path files, open it and merge it all in one text
func processFile(filePath string) (text string, err error) {
	// Read the content of the file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("TU VIEJA")
		return "", fmt.Errorf("error opening the file %s: %s", filePath, err)
	}

	return string(fileData), nil
}
