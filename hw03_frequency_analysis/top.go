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
	wordsLen := len(words)
	temp := make([]string, 0, wordsLen)
	for key := range words {
		temp = append(temp, key)
	}
	sort.Slice(temp, func(i, j int) bool {
		return (words[temp[i]] > words[temp[j]]) ||
			(words[temp[i]] == words[temp[j]] && temp[i] < temp[j])
	})
	if wordsLen > count {
		temp = temp[0:count]
	}
	return temp
}
