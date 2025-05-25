# 🌐 xlorem

**`xlorem`** is a highly configurable and grammatically structured Lorem Ipsum generator written in Go. It supports generating **paragraphs**, **sentences**, and **words** — each constructed with Latin-style noun-verb-adjective-preposition patterns for realistic placeholder content.

> 💡 Perfect for developers who are tired of boring "lorem ipsum dolor" and want something grammatically richer and more versatile.

---

## ✨ Features

- ✅ Paragraph, sentence, and word generation modes
- ✅ Proper grammatical structure using word classes (noun, verb, etc.)
- ✅ Optional `Lorem ipsum dolor sit amet.` prefix
- ✅ Configurable word/sentence ranges for realistic variation
- ✅ HTML or plain-text output for paragraphs
- ✅ CLI tool for command-line use

---

## 📦 Installation

```bash
go get github.com/RohitPModani/xlorem
````

---

## 🧠 API Usage

### `Paragraphs`

```go
func Paragraphs(count int, startsWithLorem, asHTML bool, sentenceRange, wordRange []int) string
```

| Parameter         | Type              | Description                                                            |
| ----------------- | ----------------- | ---------------------------------------------------------------------- |
| `count`           | `int`             | Number of paragraphs to generate (default: `1` if `<= 0`)              |
| `startsWithLorem` | `bool`            | If `true`, prepends first paragraph with "Lorem ipsum dolor sit amet." |
| `asHTML`          | `bool`            | Wraps each paragraph in `<p>` tags if `true`                           |
| `sentenceRange`   | `[]int{min, max}` | Number of sentences per paragraph (randomly chosen within range)       |
| `wordRange`       | `[]int{min, max}` | Number of words per sentence (randomly chosen within range)            |

---

### `Sentences`

```go
func Sentences(count int, startsWithLorem bool, wordRange []int) string
```

| Parameter         | Type              | Description                                              |
| ----------------- | ----------------- | -------------------------------------------------------- |
| `count`           | `int`             | Number of sentences to generate (default: `1` if `<= 0`) |
| `startsWithLorem` | `bool`            | If `true`, prepends with "Lorem ipsum dolor sit amet."   |
| `wordRange`       | `[]int{min, max}` | Words per sentence, randomly chosen                      |

---

### `Words`

```go
func Words(count int, startsWithLorem bool) string
```

| Parameter         | Type   | Description                                                       |
| ----------------- | ------ | ----------------------------------------------------------------- |
| `count`           | `int`  | Number of words (default: `5`)                                    |
| `startsWithLorem` | `bool` | If `true` and count ≥ 5, begins with "Lorem ipsum dolor sit amet" |

---

## 🧪 Example

```go
package main

import (
	"fmt"
	"github.com/RohitPModani/xlorem"
)

func main() {
	fmt.Println(xlorem.Paragraphs(2, true, true, []int{3, 5}, []int{8, 12}))
	fmt.Println(xlorem.Sentences(3, true, []int{6, 10}))
	fmt.Println(xlorem.Words(10, true))
}
```

---

## 🖥️ CLI Usage

A simple CLI interface is also available.

```bash
go run main.go -mode=paragraphs -count=2 -sentenceRange=3-5 -wordRange=8-12 -asHTML=true
```

### Flags

| Flag               | Description                                       |
| ------------------ | ------------------------------------------------- |
| `-mode`            | One of: `paragraphs`, `sentences`, `words`        |
| `-count`           | Number of items to generate                       |
| `-startsWithLorem` | (bool) Prepend with "Lorem ipsum dolor sit amet." |
| `-asHTML`          | (bool) Only applies to `paragraphs` mode          |
| `-sentenceRange`   | Sentence count range (e.g., `3-6`)                |
| `-wordRange`       | Word count range per sentence (e.g., `8-12`)      |

---

## 📚 Word Sources

All words are pseudo-Latin or adapted from classic Lorem Ipsum and other Latin-like constructs. They're organized by:

* **Nouns**
* **Verbs**
* **Adjectives**
* **Conjunctions**
* **Prepositions**

This results in grammatically coherent outputs that sound more natural and varied.

---

## 🧼 Formatting Details

* Sentences start with a capital letter and end with a period
* HTML output wraps paragraphs with `<p>`
* No external dependencies

---

## 🚧 Roadmap Ideas

* Add support for headings and lists
* Add Markdown support for CLI
* Language-themed Lorem generators (e.g., "Startup Ipsum", "Developer Ipsum")

---

## 🤝 Contributing

PRs and ideas welcome! Fork the repo and go wild.

---

## 📜 License

MIT License — free for personal and commercial use.

---

## 🌍 Repo

**GitHub:** [github.com/RohitPModani/xlorem](https://github.com/RohitPModani/xlorem)
