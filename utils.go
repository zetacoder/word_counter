package main

import (
	"regexp"
	"sort"
	"strings"
)

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

func countTripleWordRepeats(text string) map[string]int {
	words := strings.Fields(text)
	if len(words) < 3 {
		return nil // Not enough words to form triples
	}

	tripleCounts := make(map[string]int)

	for i := 0; i < len(words)-2; i++ {
		triple := words[i : i+3]
		tripleStr := strings.Join(triple, " ")
		tripleCounts[tripleStr]++
	}

	return tripleCounts
}

type keyValue struct {
	Key   string
	Value int
}

func sortRepeatedWords(m map[string]int) []keyValue {
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
