package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

var executable string = "word_counter"

func main() {

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
					log.Println(err)
					return
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
			fullText += text
		}

		cleanedText, _ := cleanText(fullText)

		sequences := countTripleWordRepeats(cleanedText)
		sortedSeq := sortRepeatedWords(sequences)

		var top100 []keyValue
		if len(sortedSeq) >= 100 {
			top100 = sortedSeq[:100]
		} else {
			top100 = sortedSeq[:len(sortedSeq)-1]
		}

		// Print the top 100 elements
		for _, kv := range top100 {
			fmt.Printf("%d - %s\n", kv.Value, kv.Key)
		}

	} else {
		fullText := processFromStdin()
		cleanedText, _ := cleanText(fullText)

		sequences := countTripleWordRepeats(cleanedText)
		sortedSeq := sortRepeatedWords(sequences)

		var top100 []keyValue
		if len(sortedSeq) >= 100 {
			top100 = sortedSeq[:100]
		} else {
			top100 = sortedSeq[:len(sortedSeq)-1]
		}

		// Print the top 100 elements
		for _, kv := range top100 {
			fmt.Printf("%d - %s\n", kv.Value, kv.Key)
		}
	}

}
