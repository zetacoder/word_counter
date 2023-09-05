package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
)

// We have two sources:
// Arguments as file paths --> command line like: ./solution.rb file1.txt file2.txt ...)
// The program also accepts input on stdin (e.g. cat file1.txt | ./solution.rb).

// Will contain the num of times three words are repeated
var threeWords map[string]int

func main() {

	// Create a channel to receive the result from the goroutine
	resultChan := make(chan string, 1)

	// Create a WaitGroup to wait for the goroutine to finish
	var wg sync.WaitGroup

	// Start the goroutine to process file paths
	wg.Add(1)
	go func() {
		defer wg.Done()

		text, err := processFilePaths(os.Args[1:])
		if err != nil {
			log.Fatal(err)
		}
		resultChan <- text
	}()

	// In the main goroutine, you can continue with other work
	// or wait for the result from the goroutine
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect the result from the goroutine
	mergedText := <-resultChan

	// Now you can work with the mergedText variable
	fmt.Println(mergedText)

}

// processFilePaths takes as argument different path files, open it and merge it all in one text
func processFilePaths(filePaths []string) (text string, err error) {
	// Check if unsufficient arguments
	if len(filePaths) < 1 {
		return "", fmt.Errorf("unsufficient arguments or malformed: %s", fmt.Sprintf("%s", filePaths))
	}

	// Read all files and merge it into one text
	for _, path := range filePaths {
		file, err := os.ReadFile(path)
		if err != nil {
			return "", fmt.Errorf("error opening the file: %s ", err)
		}
		text += " " + string(file)
	}

	text = strings.TrimSpace(text)

	return text, nil
}

// mergeTexts makes an unique full text for both of the sources
func mergeTextsOfAllSources(texts ...string) (string, error) {
	var merged string
	for _, text := range texts {
		merged += " " + text
	}

	return merged, nil
}

// cleanText takes the full text generated from all files of some input source and clean it
// removing unnecesary characters, punctuation and endlines.
func cleanText(text string) (words []string, err error) {
	// convert all text to lower case
	text = strings.ToLower(text)

	// Compile the regular expression
	wordPattern := `\b[\p{L}\p{Nd}'’]+\b`
	re := regexp.MustCompile(wordPattern)

	// Find all matches in the text
	words = re.FindAllString(text, -1)

	return words, nil
}
