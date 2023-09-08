package main

import (
	"testing"
)

func TestNoTextsFromArguments(t *testing.T) {
	args := ""
	_, err := getInputArguments(args)
	if err == nil {
		t.Errorf("Expected error, but not found")
	}
}

func TestOneTextFromArguments(t *testing.T) {
	arg := "./texts/clean_code.txt"
	text, err := getInputArguments(arg)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	expected := []keyValue{
		{"mies van der", 3},
		{"act of repair", 2},
		{"and more so", 2},
		{"attentiveness to detail", 2},
		{"bob and i", 2},
		{"even in the", 2},
		{"focus is on", 2},
		{"in the details", 2},
		{"is faithful in", 2},
		{"keep the workplace", 2},
		{"more so of", 2},
		{"of what we", 2},
		{"one of the", 2},
		{"practice in the", 2},
		{"small things matter", 2},
		{"software development and", 2},
		{"that they are", 2},
		{"the charm of", 2},
		{"van der rohe", 2},
		{"000 mile oil", 1},
		{"10 000 mile", 1},
		{"17 years ago", 1},
		{"1951 a quality", 1},
		{"1960 is to", 1},
		{"1990s notions of", 1},
		{"19th century french", 1},
		{"50 years ago", 1},
		{"5s is a", 1},
		{"5s philosophy comprises", 1},
		{"5s principles 5s", 1},
		{"5s principles are", 1},
		{"80 or more", 1},
		{"a bad penny", 1},
		{"a book about", 1},
		{"a consistent coding", 1},
		{"a day keeps", 1},
		{"a desire for", 1},
		{"a first born", 1},
		{"a good omen", 1},
		{"a grand whole", 1},
		{"a house but", 1},
		{"a house is", 1},
		{"a level inspect", 1},
		{"a messy desk", 1},
		{"a million selfless", 1},
		{"a new room", 1},
		{"a perfect complement", 1},
		{"a piece of", 1},
		{"a place for", 1},
		{"a poem is", 1},
		{"a pound of", 1},
		{"a pristine building", 1},
		{"a product as", 1},
		{"a product backlog", 1},
		{"a profession that", 1},
		{"a project the", 1},
		{"a quality approach", 1},
		{"a remedy for", 1},
		{"a set of", 1},
		{"a small local", 1},
		{"a small thing", 1},
		{"a stitch in", 1},
		{"a twopack of", 1},
		{"a variable using", 1},
		{"abandon our code", 1},
		{"abandonment such preoccupation", 1},
		{"about 1951 a", 1},
		{"about being eager", 1},
		{"about calibrating the", 1},
		{"about doing about", 1},
		{"about having a", 1},
		{"about how to", 1},
		{"about humble concerns", 1},
		{"about integrating simple", 1},
		{"about littering your", 1},
		{"about pushing the", 1},
		{"about seiso cleanliness", 1},
		{"about shutsuke in", 1},
		{"about still architecture", 1},
		{"about that in", 1},
		{"about the role", 1},
		{"about this approach", 1},
		{"above flowed from", 1},
		{"academia humbles itself", 1},
		{"accept for granted", 1},
		{"acorns grow or", 1},
		{"act of design", 1},
		{"actions of those", 1},
		{"acts are simple", 1},
		{"acts of care", 1},
		{"add a new", 1},
		{"add through the", 1},
		{"admonishes us we", 1},
		{"admonishments the seiton", 1},
		{"advises us that", 1},
		{"agile the focus", 1},
		{"agile world bob", 1},
		{"ago build machines", 1},
		{"ago such style", 1},
		{"ago surrendered to", 1},
	}

	text, err = cleanText(text)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	sequences := getRepeatedSequences(text, numberOfConsecutivesWords)
	sortedSequences, err := sortSequences(sequences)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// Check if the sorted sequences match the expected values in terms of key and value
	for i, exp := range expected {
		if exp != sortedSequences[i] {
			t.Errorf("Mismatch at index %d: Expected %v, but got %v", i, exp, sortedSequences[i])
		}
	}
}

func TestMultipleTextsFromArguments(t *testing.T) {
	args := []string{"./texts/clean_code.txt", "./texts/dogs_and_cats.txt", "./texts/golang_is.txt", "./texts/gutenberg.txt"}
	var texts string
	for _, filePath := range args {
		text, err := getInputArguments(filePath)
		if err != nil {
			t.Errorf("Error: %s", err)
			return
		}
		texts += text + " "
	}

	texts, err := cleanText(texts)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	expected := []keyValue{
		{Key: "of the same", Value: 320},
		{Key: "the same species", Value: 130},
		{Key: "conditions of life", Value: 125},
		{Key: "in the same", Value: 117},
		{Key: "of natural selection", Value: 115},
		{Key: "from each other", Value: 104},
		{Key: "species of the", Value: 101},
		{Key: "on the other", Value: 90},
		{Key: "the other hand", Value: 82},
		{Key: "the case of", Value: 78},
		{Key: "the theory of", Value: 76},
		{Key: "of the world", Value: 74},
		{Key: "parts of the", Value: 74},
		{Key: "some of the", Value: 74},
		{Key: "through natural selection", Value: 70},
		{Key: "with respect to", Value: 69},
		{Key: "in the case", Value: 67},
		{Key: "it may be", Value: 65},
		{Key: "the inhabitants of", Value: 65},
		{Key: "the species of", Value: 65},
		{Key: "of the species", Value: 64},
		{Key: "the same genus", Value: 63},
		{Key: "that of the", Value: 62},
		{Key: "forms of life", Value: 61},
		{Key: "individuals of the", Value: 59},
		{Key: "as far as", Value: 58},
		{Key: "those of the", Value: 57},
		{Key: "part of the", Value: 56},
		{Key: "the number of", Value: 56},
		{Key: "the principle of", Value: 54},
		{Key: "in this case", Value: 53},
		{Key: "the nature of", Value: 53},
		{Key: "nature of the", Value: 52},
		{Key: "to each other", Value: 52},
		{Key: "golang is concurrent", Value: 51},
		{Key: "golang is fast", Value: 51},
		{Key: "on the same", Value: 51},
		{Key: "as in the", Value: 50},
	}

	sequences := getRepeatedSequences(texts, numberOfConsecutivesWords)
	sortedSequences, err := sortSequences(sequences)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// we check if sortedSequences are equal to expected
	for i, exp := range expected {
		if exp != sortedSequences[i] {
			t.Errorf("Mismatch at index %d: Expected %v, but got %v", i, exp, sortedSequences[i])
		}
	}

}

func TestEmptyText(t *testing.T) {
	text := ""
	_, err := cleanText(text)
	if err == nil {
		t.Errorf("Expected error, but not found")
	}
}

func TestOneTextWithSameValues(t *testing.T) {
	text := "The quick brOwn Fox Jumps !!over the. lazy dog"
	text, err := cleanText(text)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	sequences := getRepeatedSequences(text, numberOfConsecutivesWords)
	sortedSequences, err := sortSequences(sequences)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	expected := []keyValue{
		{Value: 1, Key: "brown fox jumps"},
		{Value: 1, Key: "fox jumps over"},
		{Value: 1, Key: "jumps over the"},
		{Value: 1, Key: "over the lazy"},
		{Value: 1, Key: "quick brown fox"},
		{Value: 1, Key: "the lazy dog"},
	}

	// we check if sortedSequences are equal to expected
	for i, exp := range expected {
		if exp != sortedSequences[i] {
			t.Errorf("Mismatch at index %d: Expected %v, but got %v", i, exp, sortedSequences[i])
		}
	}

}
