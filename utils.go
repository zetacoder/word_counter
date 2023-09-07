package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// Takes the full text generated from all files of some input source and clean it removing unnecesary characters, punctuation and endlines.
func cleanText(text string) (cleanedText string, err error) {
	// convert all text to lower case
	text = strings.ToLower(text)

	// Compile the regular expression
	wordPattern := `\b[\p{L}\p{Nd}'’]+\b`
	re := regexp.MustCompile(wordPattern)

	// Find all matches in the text
	words := re.FindAllString(text, -1)

	cleanedText = strings.Join(words, " ")

	return cleanedText, nil
}

// Gets how many times N consecutives words are repeated within a text
func getRepeatedSequences(text string, numOfWords int) map[string]int {
	words := strings.Fields(text)
	if len(words) < numOfWords {
		return nil // Not enough words to form the word patterns
	}

	repeatedSeqs := make(map[string]int)

	for i := 0; i < len(words)-(numOfWords-1); i++ {
		sequence := words[i : i+numOfWords]
		seqStr := strings.Join(sequence, " ")
		repeatedSeqs[seqStr]++
	}

	return repeatedSeqs
}

type keyValue struct {
	Key   string
	Value int
}

// Order the repeated sequences in descending order
func sortSequences(m map[string]int) []keyValue {
	// Convert the map to a slice of key-value pairs
	var keyValueSlice []keyValue
	for key, value := range m {
		keyValueSlice = append(keyValueSlice, keyValue{Key: key, Value: value})
	}

	// Define a custom sorting function
	sort.Slice(keyValueSlice, func(i, j int) bool {
		// Sort in descending order based on the int values
		return keyValueSlice[i].Value > keyValueSlice[j].Value
	})

	return keyValueSlice
}

func displayMostRepeatedWords(sortedSequences []keyValue, largeOfListToDisplay int) {
	var topWords []keyValue
	if len(sortedSequences) >= largeOfListToDisplay {
		topWords = sortedSequences[:largeOfListToDisplay]
	} else {
		topWords = sortedSequences[:len(sortedSequences)-1]
	}

	// Print the top N repeated sequences of words
	for _, kv := range topWords {
		fmt.Printf("%d - %s\n", kv.Value, kv.Key)
	}
}
