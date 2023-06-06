package main

import "github.com/phasewalk1/lexer"

func main() {
	tree := lexer.MakeTree()
	var tok rune = 'a'
	tree.Insert(tok)
}