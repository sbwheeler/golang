package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func wordCount(s string) map[string]int {
	wordMap := make(map[string]int)

	for _, f := range strings.Fields(s) {
		wordMap[f]++
	}
	return wordMap
}

func main() {
	wc.Test(wordCount)
}
