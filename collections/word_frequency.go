package main

import (
	"fmt"
	"strings"
)

type passage string

const PUNCTUATIONS = `.;"-`

func (text passage) wordFrequency() map[string]int {
	textWords := strings.Fields(string(text))
	wordFrequency := make(map[string]int)
	for _, word := range textWords {
		wordLowered := strings.ToLower(word)
		trimmedWord := strings.Trim(wordLowered, PUNCTUATIONS)
		wordFrequency[trimmedWord]++
	}
	return wordFrequency
}

func main() {
	text := passage(`As far as eye could reach he saw nothing but the stems of the great plants about him
receding in the violet shade, and far overhead the multiple transparency of huge leaves
filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever
he felt able he ran again; the ground continued soft and springy, covered with the same
resilient weed which was the first thing his hands had touched in Malacandra. Once or
twice a small red creature scuttled across his path, but otherwise there seemed to be no
life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned
and alone in a forest of unknown vegetation thousands or millions of miles beyond the
reach or knowledge of man.
—C.S. Lewis, Out of the Silent Planet,
(see mng.bz/V7nO)`)
	freq := text.wordFrequency()
	fmt.Println("Printing frequency by word")
	fmt.Println("Total unique words :", len(freq))
	for word, occurence := range freq {
		fmt.Printf("%-10v:%-2v\n", word, occurence)
	}
}
