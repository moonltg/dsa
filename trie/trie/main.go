package main

import (
	"fmt"
)

type Trie struct {
	root *Node
}

type Node struct {
	children map[rune]*Node
	value    rune
	terminal bool
}

func (t *Trie) contains(word string) bool {
	node := t.root
	if node == nil || len(word) == 0 {
		return false
	}

	var ok bool
	for _, s := range word {
		node, ok = node.children[s]
		if !ok {
			break
		}
	}

	return node != nil && node.terminal
}

func (t *Trie) add(word string) {
	if len(word) == 0 {
		return
	}

	node, ok := t.root, false
	for _, s := range word {
		if t.root == nil {
			node = &Node{children: make(map[rune]*Node)}
			t.root = node
		}

		_, ok = node.children[s]
		if !ok {
			node.children[s] = &Node{value: s, children: make(map[rune]*Node)}
		}

		node = node.children[s]
	}

	node.terminal = true
}

func main() {
	fmt.Println("Trie!")

	trie := &Trie{}

	words := []string{"one", "only", "tuesday", "twice", "two", "twenty", "日本語", "日本人"}

	for _, word := range words {
		trie.add(word)
	}

	words = []string{"one", "only", "tuesday", "twice", "日本語", "日本人", "日", "two", "twenty", "three", "four"}
	for _, word := range words {
		fmt.Printf("%s: %v\n", word, trie.contains(word))
	}
}
