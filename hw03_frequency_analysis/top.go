package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(s string) []string {
	if len(s) == 0 {
		return nil
	}
	words := make(map[string]int)
	strSlice := strings.Fields(s)
	for _, word := range strSlice {
		if len(word) > 0 {
			words[word]++
		}
	}
	return sortWords(10, words)
}

func sortWords(count int, words map[string]int) []string {
	type keyValue struct {
		Key   string
		Value int
	}
	wordsLen := len(words)
	sortedStruct := make([]keyValue, 0, wordsLen)

	for key, value := range words {
		sortedStruct = append(sortedStruct, keyValue{key, value})
	}

	sort.Slice(sortedStruct, func(i, j int) bool {
		return (sortedStruct[i].Value > sortedStruct[j].Value) ||
			(sortedStruct[i].Value == sortedStruct[j].Value && sortedStruct[i].Key < sortedStruct[j].Key)
	})

	result := make([]string, 0, count)

	for i := 0; (i < count) && (i < wordsLen); i++ {
		result = append(result, sortedStruct[i].Key)
	}

	return result
}
