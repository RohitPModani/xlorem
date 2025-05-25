package xlorem

import (
	"strings"
	"testing"
	"unicode"
)

func TestParagraphs(t *testing.T) {
	tests := []struct {
		name           string
		count          int
		startsWithLorem bool
		asHTML         bool
		sentenceRange  []int
		wordRange      []int
	}{
		{"Basic paragraph", 1, false, false, []int{3, 5}, []int{7, 12}},
		{"Multiple paragraphs", 3, true, false, []int{4, 6}, []int{5, 10}},
		{"HTML paragraphs", 2, false, true, []int{2, 4}, []int{8, 15}},
		{"Invalid count", 0, false, false, []int{3, 5}, []int{7, 12}},
		{"Custom ranges", 1, true, false, []int{1, 3}, []int{3, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Paragraphs(tt.count, tt.startsWithLorem, tt.asHTML, tt.sentenceRange, tt.wordRange)

			if tt.count <= 0 {
				if result == "" {
					t.Error("Expected non-empty result for zero count")
				}
				return
			}

			if tt.asHTML {
				if !strings.HasPrefix(result, "<p>") || !strings.Contains(result, "</p>") {
					t.Error("HTML output missing p tags")
				}
			}

			if tt.startsWithLorem {
				if !strings.HasPrefix(result, "Lorem ipsum") {
					t.Error("Paragraph should start with Lorem ipsum")
				}
			}

			// Verify sentence count in first paragraph
			firstPara := result
			if tt.asHTML {
				parts := strings.Split(result, "<p>")
				if len(parts) < 2 {
					t.Fatal("Invalid HTML structure")
				}
				firstPara = strings.Split(parts[1], "</p>")[0]
			}
			sentences := strings.Split(firstPara, ". ")
			minSentences := tt.sentenceRange[0]
			if len(sentences) < minSentences {
				t.Errorf("Expected at least %d sentences, got %d", minSentences, len(sentences))
			}
		})
	}
}

func TestSentences(t *testing.T) {
	tests := []struct {
		name           string
		count          int
		startsWithLorem bool
		wordRange      []int
		expectedMin    int
	}{
		{"Single sentence", 1, false, nil, 1},
		{"Multiple sentences", 3, true, []int{5, 8}, 3},
		{"Custom word range", 2, false, []int{10, 15}, 2},
		{"Zero count", 0, false, nil, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sentences(tt.count, tt.startsWithLorem, tt.wordRange)

			sentences := strings.Split(result, ". ")
			// Account for possible empty last element
			if len(sentences) > 0 && sentences[len(sentences)-1] == "" {
				sentences = sentences[:len(sentences)-1]
			}

			if len(sentences) < tt.expectedMin {
				t.Errorf("Expected at least %d sentences, got %d", tt.expectedMin, len(sentences))
			}

			if tt.startsWithLorem && !strings.HasPrefix(result, "Lorem ipsum") {
				t.Error("Should start with Lorem ipsum")
			}

			// Test sentence capitalization
			for _, sent := range sentences {
				if len(sent) > 0 {
					firstChar := []rune(sent)[0]
					if !unicode.IsUpper(firstChar) {
						t.Errorf("Sentence should start with uppercase: %q", sent)
					}
				}
			}
		})
	}
}

func TestGenerateGrammaticalSentence(t *testing.T) {
	tests := []struct {
		name      string
		wordCount int
		minWords  int
	}{
		{"Short sentence", 5, 4},
		{"Medium sentence", 10, 8},
		{"Long sentence", 20, 15},
		{"Minimum words", 4, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateGrammaticalSentence(tt.wordCount)

			words := strings.Fields(result)
			if len(words) < tt.minWords {
				t.Errorf("Sentence too short: got %d words, want at least %d: %q", 
					len(words), tt.minWords, result)
			}

			// Check capitalization
			if !unicode.IsUpper([]rune(result)[0]) {
				t.Errorf("Sentence should start with uppercase: %q", result)
			}

			// Check ending punctuation
			if !strings.HasSuffix(result, ".") {
				t.Errorf("Sentence should end with period: %q", result)
			}

			// Verify grammatical structure (at least noun + verb)
			hasNoun := false
			hasVerb := false
			for _, word := range words {
				if contains(nouns, strings.ToLower(word)) {
					hasNoun = true
				}
				if contains(verbs, strings.ToLower(word)) {
					hasVerb = true
				}
			}
			if !hasNoun || !hasVerb {
				t.Errorf("Sentence lacks basic structure (noun + verb): %q", result)
			}
		})
	}
}

func TestWords(t *testing.T) {
	tests := []struct {
		name           string
		count          int
		startsWithLorem bool
	}{
		{"Basic word list", 5, false},
		{"Lorem prefix", 10, true},
		{"Single word", 1, false},
		{"Zero count", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Words(tt.count, tt.startsWithLorem)

			words := strings.Split(result, " ")
			if tt.count <= 0 && len(words) != defaultWordCount {
				t.Errorf("Expected default word count for zero input")
			}

			if tt.startsWithLorem && tt.count >= len(loremPrefix) {
				firstWords := strings.Join(words[:5], " ")
				if firstWords != "Lorem ipsum dolor sit amet" {
					t.Errorf("Expected Lorem ipsum prefix, got %q", firstWords)
				}
			}

			if tt.count > 0 && len(words) != tt.count {
				t.Errorf("Expected %d words, got %d", tt.count, len(words))
			}
		})
	}
}

func TestRandomWordFunctions(t *testing.T) {
	testCases := []struct {
		name string
		fn   func() string
		list []string
	}{
		{"randomNoun", randomNoun, nouns},
		{"randomVerb", randomVerb, verbs},
		{"randomAdjective", randomAdjective, adjectives},
		{"randomConjunction", randomConjunction, conjunctions},
		{"randomPreposition", randomPreposition, prepositions},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test multiple runs to ensure randomness
			found := make(map[string]bool)
			for i := 0; i < 100; i++ {
				word := tc.fn()
				if !contains(tc.list, word) {
					t.Errorf("%s returned word not in list: %s", tc.name, word)
				}
				found[word] = true
			}

			// Verify we're getting different words (not testing randomness strictly,
			// just that we're not getting the same word every time)
			if len(found) < 3 {
				t.Errorf("%s seems to not be random enough, only got %d unique words", tc.name, len(found))
			}
		})
	}
}

func TestCapitalize(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test", "Test"},
		{"TEST", "TEST"},
		{"123test", "123test"},
		{"", ""},
		{"ñame", "Ñame"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := capitalize(tt.input)
			if result != tt.expected {
				t.Errorf("capitalize(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func contains(list []string, item string) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}