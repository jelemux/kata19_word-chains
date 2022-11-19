package dict

import (
	"embed"
	"fmt"
	"strings"

	"github.com/jelemux/kata19_word-chains/path"
)

//go:embed wordlist.txt
var wordlist string
var _ = embed.FS{}

type Dictionary struct {
	words []string
}

func NewFromEmbedded() *Dictionary {
	return NewFromRaw(wordlist)
}

func NewFromRaw(wordList string) *Dictionary {
	return New(strings.Split(wordList, "\n"))
}

func New(words []string) *Dictionary {
	return &Dictionary{words}
}

func (d *Dictionary) ConnectWords(word1, word2 string) ([]string, error) {
	err := validate(word1, word2)
	if err != nil {
		return nil, err
	}

	filteredDict := d.filterByLength(len(word1))

	resultChan := make(chan path.Path, 1)
	finishedChan := make(chan bool, 1)

	go filteredDict.walk(path.New(), word1, word2, resultChan, finishedChan)

	select {
	case result := <-resultChan:
		return result.Steps, nil
	case <-finishedChan:
		close(finishedChan)
		return nil, fmt.Errorf("failed to find path from '%s' to '%s'", word1, word2)
	}
}

func (d *Dictionary) walk(traversed path.Path, currentWord string, targetWord string, result chan path.Path, finished chan bool) {
	if !(traversed.Contains(currentWord)) {

		newPath := traversed.Add(currentWord)
		differByOne := d.findDifferByOneLetter(currentWord)

		numberOfChildren := len(differByOne)

		if numberOfChildren > 0 {
			finishedChildren := make(chan bool, numberOfChildren)

			for _, nextWord := range differByOne {
				if strings.EqualFold(nextWord, targetWord) {
					result <- newPath.Add(targetWord)
					finishedChildren <- true
					continue
				}

				go d.walk(newPath.Clone(), nextWord, targetWord, result, finishedChildren)
			}

			// wait for children
			for range finishedChildren {}
			close(finishedChildren)
		}
	}

	finished <- true
}

func (d *Dictionary) findDifferByOneLetter(originalWord string) []string {
	var differByOne []string
	for _, word := range d.words {
		if differsByOneLetter(originalWord, word) {
			differByOne = append(differByOne, word)
		}
	}

	return differByOne
}

func (d *Dictionary) filterByLength(length int) *Dictionary {
	var filtered []string
	for _, word := range d.words {
		if len(word) == length {
			filtered = append(filtered, word)
		}
	}

	return New(filtered)
}

func differsByOneLetter(word1, word2 string) bool {
	differentLetters := 0
	for i := 0; i < len(word1); i++ {
		if strings.ToLower(word1)[i] != strings.ToLower(word2)[i] {
			differentLetters++
		}
		if differentLetters > 1 {
			return false
		}
	}

	return differentLetters != 0
}

func validate(word1, word2 string) error {
	if len(word1) != len(word2) {
		return fmt.Errorf("words '%s' and '%s' must be of same length", word1, word2)
	}

	return nil
}