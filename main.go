package main

import (
	"os"

	"github.com/jelemux/kata19_word-chains/dict"
)

func main() {
	dictionary := dict.NewFromEmbedded()
	steps, err := dictionary.ConnectWords("suber", "Typha")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	for _, step := range steps {
		println(step)
	}
}