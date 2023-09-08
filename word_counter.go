package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

// word_counter contains all the functions necessary to takes different sources of inputs
// and extract the N most repeated Y consecutive words, case, spaces and punctuation insensitive.

const executable string = "word_counter"
const numberOfConsecutivesWords int = 3 // You can set the numbers of word you want to use.
const largeOfListToDisplay int = 100    // You can set the large of the list of repeated words to display in output.

// runWordCounter wraps all the necesarry functions and execute the programm
func runWordCounter() error {
	// Create a channel to receive results from goroutines
	resultChan := make(chan string)

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Extract file paths from command-line arguments
	filePaths := os.Args[1:]

	// Check if command-line arguments are provided and process the inputs
	if areCommandLineArguments(filePaths, os.Args[0], executable) {
		// Start a goroutine for each file
		for _, path := range filePaths {
			wg.Add(1)
			go func(filePath string) {
				defer wg.Done()
				text, err := getInputArguments(filePath)
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
		texts := ""
		for text := range resultChan {
			texts += text + " "
		}

		// Clean the text, get the repeated words and display it in output.
		err := process(texts)
		if err != nil {
			return err
		}

		return nil
	}

	// get input from stdin. eg. cat arg1 arg2 ... argN | ./executable
	texts := getInputStdin()

	// Clean the text, get the repeated words and display it in output.
	err := process(texts)
	if err != nil {
		return err
	}

	return nil
}

// Process the input, get repeated words, sort and displays it
func process(texts string) error {
	text, err := cleanText(texts)
	if err != nil {
		return err
	}
	sequences := getRepeatedSequences(text, numberOfConsecutivesWords)
	sortedSequences, err := sortSequences(sequences)
	if err != nil {
		return err
	}
	displayMostRepeatedWords(sortedSequences, largeOfListToDisplay)

	return nil
}

// Takes as argument different path files, open it and merge it all in one text. e.g when in cli: ./executable arg1.txt arg2.txt ... argN.txt
func getInputArguments(filePath string) (text string, err error) {
	// Read the content of the file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error opening the file %s: %s", filePath, err)
	}

	return string(fileData), nil
}

// Checks if the input is passed as arguments. e.g ./executable arg1 arg2 ... argN
func areCommandLineArguments(filePaths []string, args string, executable string) bool {
	if len(filePaths) > 0 && strings.Contains(args, executable) {
		return true
	}

	return false
}

// Takes as argument the inputs from stdin and put them in pipe to be process. e.g: cat arg1.txt arg2.txt ... argN.txt | ./executable
func getInputStdin() string {
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
