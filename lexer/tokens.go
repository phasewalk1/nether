package lexer

type Token struct {
	value rune
	pos uint32
	next *Token
}

func (t *Token) Value() rune {
	return t.value
}

func (t *Token) Pos() uint32 {
	return t.pos
}

func (t *Token) Next() *Token {
	return t.next
}

func newNaiveToken(value rune) *Token {
	return &Token{value: value, pos: 0, next: nil}
}