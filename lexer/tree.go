package lexer

type Tree interface {
	Root() *Token
	Insert(value rune) uint32
}

type NetherTree struct {
	head *Token
}

func setHead (tree *NetherTree, tok *Token) {
	tree.head = tok
}

func MakeTree() *NetherTree {
	return &NetherTree{head: nil}
}

func (t NetherTree) Root() *Token {
	return t.head
}

func (t *NetherTree) Insert(value rune) uint32 {
	var pos uint32 = 0
	var prev *Token = nil
	var curr *Token = t.head

	for curr != nil {
		pos++
		prev = curr
		curr = curr.next
	}

	newToken := &Token{value: value, pos: pos, next: nil}

	if prev == nil {
		setHead(t, newToken)
	} else {
		prev.next = newToken
	}

	return pos
}