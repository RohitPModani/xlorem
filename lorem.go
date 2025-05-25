package xlorem

import (
	"math/rand"
	"strings"
	"time"
	"unicode"
)

// Word categories for grammatical structure
var (
	nouns = []string{
		"adipisicing", "amet", "anim", "aute", "cillum", "commodo", "consectetur", 
		"consequat", "culpa", "cupidatat", "deserunt", "dolor", "dolore", "elit", 
		"enim", "esse", "est", "eu", "ex", "excepteur", "exercitation", "fugiat", 
		"id", "in", "incididunt", "ipsum", "irure", "labore", "laboris", "laborum", 
		"lorem", "magna", "minim", "mollit", "nisi", "non", "nostrud", "nulla", 
		"occaecat", "officia", "pariatur", "proident", "qui", "quis", "reprehenderit", 
		"sint", "sit", "sunt", "tempor", "ullamco", "ut", "velit", "veniam", "voluptate",
	}

	verbs = []string{
		"ad", "adipiscing", "aliqua", "aliquip", "beatae", "cred", "debitis", 
		"delectus", "deleniti", "do", "duis", "ea", "eiusmod", "epicuri", "et", 
		"facere", "fuga", "graeco", "habeo", "iisque", "interpretaris", "ius", 
		"latine", "liberavisse", "mentitum", "pertinax", "philosophia", "placerat", 
		"ponderum", "populo", "pri", "quando", "referrentur", "reformidans", "regione", 
		"scaevola", "scripta", "sed", "suspendisse", "tacimates", "tantas", "tation", 
		"te", "theophrastus", "tibique", "timeam", "vim", "vivendo", "vituperata",
	}

	adjectives = []string{
		"accusamus", "accusantium", "adipisci", "alias", "aliquid", "asperiores", 
		"aspernatur", "atque", "aut", "autem", "blanditiis", "corporis", "corrupti", 
		"dexter", "distinctio", "dolorem", "doloremque", "dolores", "ducimus", 
		"eaque", "earum", "eos", "error", "eveniet", "expedita", "explicabo", 
		"facilis", "harum", "hic", "illo", "impedit", "inventore", "iusto", 
		"laboriosam", "laudantium", "libero", "maiores", "maxime", "modi", 
		"molestiae", "molestias", "nam", "natus", "necessitatibus", "nemo", 
		"neque", "nesciunt", "nobis", "odio", "odit", "omnis", "optio", 
		"perspiciatis", "placeat", "porro", "possimus", "praesentium", "quae", 
		"quam", "quas", "quasi", "quibusdam", "quidem", "quisquam", "quo", 
		"quod", "recusandae", "repellat", "repellendus", "repudiandae", "rerum", 
		"saepe", "sapiente", "sequi", "similique", "soluta", "tempora", 
		"tenetur", "totam", "ullam", "unde", "veritatis", "vitae", "voluptas", 
		"voluptatem", "voluptates", "voluptatibus", "voluptatum",
	}

	conjunctions = []string{
		"ac", "at", "cum", "de", "ergo", "etiam", "ex", "haec", "hic", "iam", 
		"igitur", "interdum", "ita", "item", "mox", "nec", "neque", "per", 
		"post", "quamquam", "quoniam", "sed", "semper", "si", "sic", "sub", 
		"sui", "tam", "tamen", "trans", "tu", "tum", "ubi", "ultro", "unde", 
		"utrum", "vero",
	}

	prepositions = []string{
		"ad", "ante", "apud", "circum", "contra", "inter", "intra", "ob", 
		"per", "post", "prae", "prope", "propter", "sub", "super", "trans",
	}

)

const (
	defaultMinSentences = 3
	defaultMaxSentences = 6
	defaultMinWords     = 8
	defaultMaxWords     = 15
	defaultWordCount    = 5
)

var (
	loremPrefix = []string{"Lorem", "ipsum", "dolor", "sit", "amet"}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Paragraphs generates grammatically structured lorem ipsum paragraphs
func Paragraphs(count int, startsWithLorem, asHTML bool, sentenceRange, wordRange []int) string {
	if count <= 0 {
		count = 1
	}

	paras := make([]string, 0, count)
	for i := 0; i < count; i++ {
		sentCount := validateRange(sentenceRange, defaultMinSentences, defaultMaxSentences)
		sentences := make([]string, 0, sentCount)
		
		for j := 0; j < sentCount; j++ {
			wordCount := validateRange(wordRange, defaultMinWords, defaultMaxWords)
			sentences = append(sentences, generateGrammaticalSentence(wordCount))
		}
		
		para := strings.Join(sentences, " ")
		paras = append(paras, para)
	}

	if startsWithLorem && len(paras) > 0 {
		paras[0] = prependLorem(paras[0])
	}

	if asHTML {
		for i := range paras {
			paras[i] = "<p>" + paras[i] + "</p>"
		}
		return strings.Join(paras, "\n")
	}
	return strings.Join(paras, "\n\n")
}

// Sentences generates grammatically correct lorem ipsum sentences
func Sentences(count int, startsWithLorem bool, wordRange []int) string {
	if count <= 0 {
		count = 1
	}

	sentences := make([]string, 0, count)
	for i := 0; i < count; i++ {
		wc := validateRange(wordRange, defaultMinWords, defaultMaxWords)
		sentences = append(sentences, generateGrammaticalSentence(wc))
	}

	if startsWithLorem && len(sentences) > 0 {
		sentences[0] = prependLorem(sentences[0])
	}

	return strings.Join(sentences, " ")
}

// Words generates lorem ipsum words with grammatical structure
func Words(count int, startsWithLorem bool) string {
	if count <= 0 {
		count = defaultWordCount
	}
	
	words := generateGrammaticalPhrase(count)
	if startsWithLorem && count >= len(loremPrefix) {
		copy(words, loremPrefix)
	}
	return strings.Join(words, " ")
}

// Helper: validateRange checks and returns a valid number within a range
func validateRange(r []int, defMin, defMax int) int {
	if len(r) == 2 && r[0] > 0 && r[1] >= r[0] {
		return rand.Intn(r[1]-r[0]+1) + r[0]
	}
	return rand.Intn(defMax-defMin+1) + defMin
}

// Helper: generateGrammaticalSentence creates a properly structured sentence
func generateGrammaticalSentence(wordCount int) string {
    words := make([]string, 0, wordCount)

    // Ensure basic sentence structure: noun + verb
    words = append(words, capitalize(randomNoun()))
    words = append(words, randomVerb())

    // Add additional words to reach wordCount (excluding period)
    remainingWords := wordCount - 2 // Account for subject and verb
    for i := 0; i < remainingWords; i++ {
        switch rand.Intn(4) {
        case 0:
            words = append(words, randomAdjective())
        case 1:
            words = append(words, randomNoun())
        case 2:
            words = append(words, randomPreposition())
        case 3:
            words = append(words, randomConjunction())
        }
    }

    // Append period to the last word
    if len(words) > 0 {
        words[len(words)-1] += "."
    } else {
        words = append(words, ".")
    }

    return strings.Join(words, " ")
}

// Helper: generateGrammaticalPhrase creates a meaningful word sequence
func generateGrammaticalPhrase(count int) []string {
	words := make([]string, 0, count)
	
	for i := 0; i < count; i++ {
		switch {
		case i == 0:
			words = append(words, capitalize(randomNoun()))
		case i == 1:
			words = append(words, randomVerb())
		default:
			switch rand.Intn(3) {
			case 0:
				words = append(words, randomAdjective())
			case 1:
				words = append(words, randomNoun())
			case 2:
				words = append(words, randomPreposition())
			}
		}
	}
	
	return words
}

// Word category helpers
func randomNoun() string {
	return nouns[rand.Intn(len(nouns))]
}

func randomVerb() string {
	return verbs[rand.Intn(len(verbs))]
}

func randomAdjective() string {
	return adjectives[rand.Intn(len(adjectives))]
}

func randomConjunction() string {
	return conjunctions[rand.Intn(len(conjunctions))]
}

func randomPreposition() string {
	return prepositions[rand.Intn(len(prepositions))]
}

// Helper: capitalize capitalizes the first letter of a word
func capitalize(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToTitle(r[0])
	return string(r)
}

// Helper: endsWithPunctuation checks if word ends with punctuation
func endsWithPunctuation(s string) bool {
	if len(s) == 0 {
		return false
	}
	lastChar := s[len(s)-1]
	return lastChar == '.' || lastChar == '!' || lastChar == '?'
}

// Helper: prependLorem adds the standard lorem prefix
func prependLorem(original string) string {
	sentence := strings.Join(loremPrefix, " ") + "."
	return sentence + " " + original
}