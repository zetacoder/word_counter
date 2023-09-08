package main

import (
	"testing"
)

func TestCleanText(t *testing.T) {
	dirtyTexts := []string{`ThE............. qUick ---!!<<>>	
	bRown foX Jumps !!over !!the @@lazy dog`, `some ..-.,<<+++ Skills !!???¡¡¿¡reQuireS
	 a lot of practice,`, `golang 
	 is a progRamming
	  lAnguage cReated at
	   ..,.,.,.,.,.,.,.,.,google 	that is concurrent, compiled, 
	   !!! garbage-collecteD anD faSt`}

	expected := []string{"the quick brown fox jumps over the lazy dog", "some skills requires a lot of practice", "golang is a programming language created at google that is concurrent compiled garbage collected and fast"}

	for i, text := range dirtyTexts {
		cleanedText, err := cleanText(text)
		if err != nil {
			t.Errorf("Error: %s", err)
		}

		if cleanedText != expected[i] {
			t.Errorf("Expected %s, but got %s", expected[i], cleanedText)
		}
	}
}

func TestGetRepeatedSequences(t *testing.T) {
	text1 := "Backend .-.-.-.-. is cool! backend is cOOl! backend is cool! backend is cool! backend -.,-.,-.,-.,!!)) is cool!"
	text2 := "sometimes Is useful to be persistent in studying and practicing even if it is hard. SomeTimes is useful to learn from others!"
	text3 := ""

	combined := text1 + " " + text2 + " " + text3
	text, err := cleanText(combined)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	sequences := getRepeatedSequences(text, numberOfConsecutivesWords)
	sortedSequences, err := sortSequences(sequences)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	expected := []keyValue{
		{Value: 5, Key: "backend is cool"},
		{Value: 2, Key: "is useful to"},
		{Value: 2, Key: "sometimes is useful"},
	}

	// we check if both the key and value are in the list
	for _, expected := range expected {
		found := false
		for _, result := range sortedSequences {
			if expected.Key == result.Key && expected.Value == result.Value {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected %s, but not found", expected.Key)
		}
	}

}

func TestSortSequence(t *testing.T) {
	sequenceCounter := map[string]int{
		"sometimes is useful":   2,
		"backend is fast":       1,
		"backend is concurrent": 1,
		"backend is compiled":   1,
		"backend is cool":       5,
		"is useful to":          2,
	}

	sortedSequences, err := sortSequences(sequenceCounter)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	expected := []keyValue{
		{Value: 5, Key: "backend is cool"},
		{Value: 2, Key: "is useful to"},
		{Value: 2, Key: "sometimes is useful"},
		{Value: 1, Key: "backend is compiled"},
		{Value: 1, Key: "backend is concurrent"},
		{Value: 1, Key: "backend is fast"},
	}

	// Check if the sorted sequences match the expected values in terms of key, value, and order
	for i, exp := range expected {
		if exp.Key != sortedSequences[i].Key || exp.Value != sortedSequences[i].Value {
			t.Errorf("Mismatch at index %d: Expected %s (value: %d), but got %s (value: %d)", i, exp.Key, exp.Value, sortedSequences[i].Key, sortedSequences[i].Value)
		}
	}

	sequenceCounter2 := map[string]int{
		"as we navigate":                       2,
		"propels us forward":                   2,
		"acts of kindness":                     2,
		"small acts of":                        2,
		"the beauty of":                        2,
		"things that matter":                   2,
		"of our lives":                         2,
		"small things that":                    2,
		"the essence of":                       2,
		"the charm of":                         2,
		"that propels us":                      2,
		"of the human":                         2,
		"of life life":                         2,
		"that encapsulates the":                1,
		"a well-maintained":                    1,
		"to progress progress":                 1,
		"human experience where":               1,
		"change a constant":                    1,
		"the grand scheme":                     1,
		"process lies in":                      1,
		"the small things":                     1,
		"navigate this journey":                1,
		"of creativity where":                  1,
		"unfolds in every":                     1,
		"ever-evolving journey":                1,
		"embrace new perspectives":             1,
		"the highway of":                       1,
		"software that delivers":               1,
		"new horizons and":                     1,
		"of life itself":                       1,
		"element plays a":                      1,
		"essence of life's":                    1,
		"of nature and":                        1,
		"to the grand":                         1,
		"that grace our":                       1,
		"in the delicate":                      1,
		"a pristine building":                  1,
		"impact the ripple":                    1,
		"realm of creativity":                  1,
		"uphold the highest":                   1,
		"kindness can have":                    1,
		"explore new horizons":                 1,
		"even the most":                        1,
		"life life at":                         1,
		"is a grand":                           1,
		"change is the":                        1,
		"intricacies reveal themselves":        1,
		"our world world":                      1,
		"in the wild":                          1,
		"of experiences emotions":              1,
		"statistically significant indicators": 1,
		"indicators of impact":                 1,
		"ourselves exploring the":              1,
		"it's a world":                         1,
		"humanity creativity is":               1,
		"power to influence":                   1,
		"we find ourselves":                    1,
		"connections that shape":               1,
		"and the diverse":                      1,
		"diverse species that":                 1,
		"balance of ecosystems":                1,
		"cycle interconnected with":            1,
		"and rewards like":                     1,
		"them we find":                         1,
		"way of life":                          1,
		"matter from the":                      1,
		"find solace in":                       1,
		"flowing water nature":                 1,
		"foundation of professionalism":        1,
		"willingness to change":                1,
		"whole composed of":                    1,
		"a willingness to":                     1,
		"well-maintained automobile":           1,
		"the ever-changing":                    1,
		"of even the":                          1,
		"to principles that":                   1,
		"to the larger":                        1,
		"creativity the wisdom":                1,
		"of kindness to":                       1,
		"a place where":                        1,
		"apply in every":                       1,
		"experiences emotions and":             1,
		"of kindness can":                      1,
		"have the power":                       1,
		"adventures in the":                    1,
		"ecosystems and in":                    1,
		"nature in all":                        1,
		"too as individuals":                   1,
		"sunrise and sunset":                   1,
		"individual like a":                    1,
		"is a brushstroke":                     1,
		"in every sunrise":                     1,
		"and the soothing":                     1,
		"a network of":                         1,
		"plays a crucial":                      1,
		"what we can":                          1,
		"future it's the":                      1,
		"ever-changing landscape":              1,
		"integrity these principles":           1,
		"existence as we":                      1,
	}

	sortedSequences2, err := sortSequences(sequenceCounter2)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	expected2 := []keyValue{
		{"acts of kindness", 2},
		{"as we navigate", 2},
		{"of life life", 2},
		{"of our lives", 2},
		{"of the human", 2},
		{"propels us forward", 2},
		{"small acts of", 2},
		{"small things that", 2},
		{"that propels us", 2},
		{"the beauty of", 2},
		{"the charm of", 2},
		{"the essence of", 2},
		{"things that matter", 2},
		{"a network of", 1},
		{"a place where", 1},
		{"a pristine building", 1},
		{"a well-maintained", 1},
		{"a willingness to", 1},
		{"adventures in the", 1},
		{"and rewards like", 1},
		{"and the diverse", 1},
		{"and the soothing", 1},
		{"apply in every", 1},
		{"balance of ecosystems", 1},
		{"change a constant", 1},
		{"change is the", 1},
		{"connections that shape", 1},
		{"creativity the wisdom", 1},
		{"cycle interconnected with", 1},
		{"diverse species that", 1},
		{"ecosystems and in", 1},
		{"element plays a", 1},
		{"embrace new perspectives", 1},
		{"essence of life's", 1},
		{"even the most", 1},
		{"ever-changing landscape", 1},
		{"ever-evolving journey", 1},
		{"existence as we", 1},
		{"experiences emotions and", 1},
		{"explore new horizons", 1},
		{"find solace in", 1},
		{"flowing water nature", 1},
		{"foundation of professionalism", 1},
		{"future it's the", 1},
		{"have the power", 1},
		{"human experience where", 1},
		{"humanity creativity is", 1},
		{"impact the ripple", 1},
		{"in every sunrise", 1},
		{"in the delicate", 1},
		{"in the wild", 1},
		{"indicators of impact", 1},
		{"individual like a", 1},
		{"integrity these principles", 1},
		{"intricacies reveal themselves", 1},
		{"is a brushstroke", 1},
		{"is a grand", 1},
		{"it's a world", 1},
		{"kindness can have", 1},
		{"life life at", 1},
		{"matter from the", 1},
		{"nature in all", 1},
		{"navigate this journey", 1},
		{"new horizons and", 1},
		{"of creativity where", 1},
		{"of even the", 1},
		{"of experiences emotions", 1},
		{"of kindness can", 1},
		{"of kindness to", 1},
		{"of life itself", 1},
		{"of nature and", 1},
		{"our world world", 1},
		{"ourselves exploring the", 1},
		{"plays a crucial", 1},
		{"power to influence", 1},
		{"process lies in", 1},
		{"realm of creativity", 1},
		{"software that delivers", 1},
		{"statistically significant indicators", 1},
		{"sunrise and sunset", 1},
		{"that encapsulates the", 1},
		{"that grace our", 1},
		{"the ever-changing", 1},
		{"the grand scheme", 1},
		{"the highway of", 1},
		{"the small things", 1},
		{"them we find", 1},
		{"to principles that", 1},
		{"to progress progress", 1},
		{"to the grand", 1},
		{"to the larger", 1},
		{"too as individuals", 1},
		{"unfolds in every", 1},
		{"uphold the highest", 1},
		{"way of life", 1},
		{"we find ourselves", 1},
		{"well-maintained automobile", 1},
		{"what we can", 1},
		{"whole composed of", 1},
		{"willingness to change", 1},
	}

	// Check if the sorted sequences match the expected values in terms of key, value, and order
	for i, exp := range expected2 {
		if exp.Key != sortedSequences2[i].Key || exp.Value != sortedSequences2[i].Value {
			t.Errorf("Mismatch at index %d: Expected %s (value: %d), but got %s (value: %d)", i, exp.Key, exp.Value, sortedSequences2[i].Key, sortedSequences2[i].Value)
		}
	}

}
