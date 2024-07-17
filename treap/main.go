package main

import (
	"math/rand"

	"github.com/mpanelo/go-dsa/treap/internal/treap"
)

var words = []string{"tree", "dog", "cat", "house", "tv", "toys", "couch", "car", "park", "mailbox", "trash"}

func main() {
	t := &treap.Treap[string]{}

	for i := 0; i < 5; i++ {
		t.Insert(randomWord())
	}

	t.InOrderPrint()
}

func randomWord() string {
	return words[rand.Intn(len(words))]
}
