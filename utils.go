package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// Utils contains functions that are used indepently of the input source to clean, count and sort the words

// Struct that holds the sequence of words and how many times it is repeated
type keyValue struct {
	Key   string
	Value int
}

// Takes the full text generated from all files of some input source and clean it removing unnecesary characters, punctuation and endlines.
func cleanText(text string) (cleanedText string, err error) {
	// convert all text to lower case
	text = strings.ToLower(text)

	// Compile the regular expression
	wordPattern := `\b[\p{L}\p{Nd}'’]+\b`
	re := regexp.MustCompile(wordPattern)

	// Find all matches in the text
	words := re.FindAllString(text, -1)

	if len(words) < numberOfConsecutivesWords {
		return "", fmt.Errorf("Not enough words to match the word patterns. Have: %d Need minimum: %d", len(words), numberOfConsecutivesWords)
	}

	cleanedText = strings.Join(words, " ")

	return cleanedText, nil
}

// Gets how many times N consecutives words are repeated within a text
func getRepeatedSequences(text string, numOfRepeatedWords int) map[string]int {
	words := strings.Fields(text)

	repeatedSeqs := make(map[string]int)

	for i := 0; i < len(words)-(numOfRepeatedWords-1); i++ {
		sequence := words[i : i+numOfRepeatedWords]
		seqStr := strings.Join(sequence, " ")
		repeatedSeqs[seqStr]++
	}

	return repeatedSeqs
}

// Sorts the sequences of words by the number of times they are repeated if the same value, order alphabetically
func sortSequences(sequenceCounter map[string]int) ([]keyValue, error) {
	if sequenceCounter == nil {
		return nil, fmt.Errorf("input map sequenceCounter is nil")
	}

	// Create a slice of key-value pairs with the original order
	var keyValueSlice []keyValue
	for key, value := range sequenceCounter {
		keyValueSlice = append(keyValueSlice, keyValue{Key: key, Value: value})
	}

	sort.SliceStable(keyValueSlice, func(i, j int) bool {
		if keyValueSlice[i].Value == keyValueSlice[j].Value {
			return keyValueSlice[i].Key < keyValueSlice[j].Key // Sort alphabetically when values are the same
		}
		return keyValueSlice[i].Value > keyValueSlice[j].Value // Sort by values in descending order
	})

	return keyValueSlice, nil
}

// displays the list of most repeated words
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
