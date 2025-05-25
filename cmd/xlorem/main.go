package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/RohitPModani/xlorem"
)

func main() {
	// Flags
	mode := flag.String("mode", "paragraphs", "Mode: paragraphs, sentences, or words")
	count := flag.Int("count", 1, "Number of paragraphs, sentences, or words")
	startsWithLorem := flag.Bool("startsWithLorem", true, "Start with 'Lorem ipsum dolor sit amet'")
	asHTML := flag.Bool("asHTML", false, "Format output as HTML (only for paragraphs)")

	// Ranges
	sentenceRange := flag.String("sentenceRange", "", "Sentence range (e.g., 3-5)")
	wordRange := flag.String("wordRange", "", "Word range (e.g., 8-12)")

	flag.Parse()

	// Parse range strings
	sentRange := parseRange(*sentenceRange)
	wordRng := parseRange(*wordRange)

	switch strings.ToLower(*mode) {
	case "paragraphs":
		result := xlorem.Paragraphs(*count, *startsWithLorem, *asHTML, sentRange, wordRng)
		fmt.Println(result)

	case "sentences":
		result := xlorem.Sentences(*count, *startsWithLorem, wordRng)
		fmt.Println(result)

	case "words":
		result := xlorem.Words(*count, *startsWithLorem)
		fmt.Println(result)

	default:
		fmt.Fprintf(os.Stderr, "Invalid mode: %s\n", *mode)
		flag.Usage()
		os.Exit(1)
	}
}

// Parses "min-max" like "3-6" into []int{3,6}
func parseRange(r string) []int {
	if r == "" {
		return nil
	}
	parts := strings.Split(r, "-")
	if len(parts) != 2 {
		return nil
	}
	var min, max int
	_, err1 := fmt.Sscanf(parts[0], "%d", &min)
	_, err2 := fmt.Sscanf(parts[1], "%d", &max)
	if err1 != nil || err2 != nil || min <= 0 || max < min {
		return nil
	}
	return []int{min, max}
}