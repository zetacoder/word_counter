package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

// wordCounter contains all the functions necessary to takes different sources of inputs
// and extract the N most repeated Y consecutive words, case, spaces and punctuation insensitive.

const executable string = "word_counter"
const numberOfConsecutivesWords int = 3  // You can set the numbers of word you want to use.
const largeOfListToDisplay int = 100 // You can set the large of the list of repeated words to display in output.

// runWordCounter wraps all the necesarry functions and execute the programm
func runWordCounter() error {
	// Create a channel to receive results from goroutines
	resultChan := make(chan string)

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Extract file paths from command-line arguments
	filePaths := os.Args[1:]

	// Check if command-line arguments are provided and the program name contains the executable name
	if len(filePaths) > 0 && strings.Contains(os.Args[0], executable) {
		// Start a goroutine for each file
		for _, path := range filePaths {
			wg.Add(1)
			go func(filePath string) {
				defer wg.Done()
				text, err := processFile(filePath)
				if err != nil {
					fmt.Printf("error processing file: %s", err)
				}
				resultChan <- text
			}(path)
		}

		// Start a goroutine to collect results and close the channel
		go func() {
			wg.Wait()
			close(resultChan)
		}()

		// Merge the text from all goroutines
		fullText := ""
		for text := range resultChan {
			fullText += text + " "
		}

		// Clean the text, get the repeated words and display it in output.
		cleanedText, _ := cleanText(fullText)
		sequences := getRepeatedSequences(cleanedText, numberOfConsecutivesWords)
		sortedSequences := sortSequences(sequences)

		displayMostRepeatedWords(sortedSequences, largeOfListToDisplay)

		return nil
	}

	// In case the input is a cat command, put the files in a pipe to be processed, c
	// clean it, get the repeated words and displays it in output.
	fullText := processFromStdin()
	text, _ := cleanText(fullText)
	sequences := getRepeatedSequences(text, numberOfConsecutivesWords)
	sortedSequences := sortSequences(sequences)

	displayMostRepeatedWords(sortedSequences, largeOfListToDisplay)

	return nil
}

// processFilePaths takes as argument different path files, open it and merge it all in one text
// e.g when in cli: ./executable arg1.txt arg2.txt ... argN.txt
func processFile(filePath string) (text string, err error) {
	// Read the content of the file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error opening the file %s: %s", filePath, err)
	}

	return string(fileData), nil
}

// processFromStdin takes as argument the inputs from stdin and put them in pipe to be process
// e.g: cat arg1.txt arg2.txt ... argN.txt | ./executable
func processFromStdin() string {
	var text string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		text += " " + line
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return text
}


