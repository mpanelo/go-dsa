package main

import (
	"fmt"
	"math/rand"

	"github.com/mpanelo/go-dsa/treap/internal/treap"
)

var words = []string{"tree", "dog", "cat", "house", "tv", "toys", "couch", "car", "park", "mailbox", "trash"}

func main() {
	t := &treap.Treap[string]{}

	for i := 0; i < 5; i++ {
		t.Insert(randomWord())
	}

	fmt.Println("Representation:")
	fmt.Println(t.Repr(true))
	fmt.Println()
	result := t.Find("tree")

	if result == nil {
		fmt.Println("I did not find 'tree'!")
	} else {
		fmt.Print("Deleting 'tree' node\n\n")
		t.Delete(result.Key)

		fmt.Println("Representation:")
		fmt.Println(t.Repr(true))
		fmt.Println()
	}
}

func randomWord() string {
	return words[rand.Intn(len(words))]
}
